package gohbase

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/jasonzzw/gohbase/filter"
	"github.com/jasonzzw/gohbase/hrpc"
	"github.com/jasonzzw/gohbase/region"
	log "github.com/sirupsen/logrus"
)

type parallelScanner struct {
	regions    []hrpc.RegionInfo
	rootClient *client
	rootScan   *hrpc.Scan
	once       sync.Once

	resultsCh   chan *hrpc.Result
	parallelism int
}

func NewParallelScanner(c *client, rpc *hrpc.Scan, parallel int) hrpc.Scanner {
	// Initialize the returning scanner.
	ret := &parallelScanner{
		regions:     []hrpc.RegionInfo{},
		rootClient:  c,
		rootScan:    rpc,
		resultsCh:   make(chan *hrpc.Result, 10000),
		parallelism: parallel,
	}

	// Scan metadata for the targeting table.
	scanRequest, err := hrpc.NewScanStr(
		context.Background(),
		"hbase:meta",
		hrpc.Filters(filter.NewPrefixFilter([]byte(rpc.Table()))))
	if err != nil {
		log.Errorf(
			"Error with scanning hbase:meta table, err=%s",
			err.Error())
		return nil
	}

	scanner := c.Scan(scanRequest)
	for {
		if scanRsp, err := scanner.Next(); err != nil {
			// Was there an error?
			if err != io.EOF {
				ret.resultsCh <- &hrpc.Result{
					Error: fmt.Errorf(
						"Failed to fetch hbase region metadata: %s",
						err.Error())}
			}

			// Abort immediately
			break

		} else if regionInfo, _, err := region.ParseRegionInfo(scanRsp); err != nil {
			// Report the error!
			ret.resultsCh <- &hrpc.Result{
				Error: fmt.Errorf(
					"Failed to parse hbase region metadata: %s",
					err.Error())}

			// Keep fetching metadata of the rest regions.

		} else {
			ret.regions = append(
				ret.regions, regionInfo)
		}
	}

	// Return the scanner!!!
	log.Infof("Retrieve %d regions for table=%s", len(ret.regions), rpc.Table())
	return ret
}

func (p *parallelScanner) fetch() {
	//put regions into the channel
	ch := make(chan hrpc.RegionInfo, 500)
	go func() {
		for _, region := range p.regions {
			ch <- region
		}
		close(ch)
	}()

	//now we start workers
	var wg sync.WaitGroup
	for i := 0; i < p.parallelism; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Retrieve a region to work on
			for region := range ch {
				scanObj, err := hrpc.NewScanRange(
					p.rootScan.Context(),
					p.rootScan.Table(),
					region.StartKey(),
					region.StopKey(),
					hrpc.Families(p.rootScan.Families()),
					hrpc.Filters(p.rootScan.Filter()))

				if err != nil {
					p.resultsCh <- &hrpc.Result{
						Error: fmt.Errorf(
							"Failed to create ScanRange: %s",
							err.Error())}

					log.Warnf("Create scan object err=%s", err)
					continue
				}

				cli := NewClient(
					p.rootClient.zkClient.GetQuorum(),
					p.rootClient.options...)

				sc := cli.Scan(scanObj)
				for {
					//scan one row
					if scanRsp, err := sc.Next(); err != nil {
						if err != io.EOF {
							p.resultsCh <- &hrpc.Result{
								Error: fmt.Errorf(
									"Failed to scan: %s",
									err.Error())}
						}

						// We've finished dealing with this region (maybe not gracefully)
						break

					} else {
						p.resultsCh <- scanRsp
					}
				}

				sc.Close()
				cli.Close()
			}
		}()
	}

	// Tear down
	wg.Wait()
	close(p.resultsCh)
}

func (p *parallelScanner) Next() (*hrpc.Result, error) {
	p.once.Do(func() {
		go p.fetch()
	})

	//return nil, nil
	if result, ok := <-p.resultsCh; ok {
		return result, result.Error
	} else {
		return nil, io.EOF
	}
}

func (p *parallelScanner) Close() error {
	return nil
}
