package worker

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"web_scraper_v2/configs"
	"web_scraper_v2/endpoint"
	"web_scraper_v2/internals/fetcher"
	jsonhandler "web_scraper_v2/internals/jsonHandler"
)

func ScraperPool(){
	
	var wg sync.WaitGroup
	workers_count := 100
	jobs := make(chan string, workers_count)

	f, err := os.Open(configs.UrlPath + "/cleaned_urls.txt")
	br := bufio.NewScanner(f)
	br.Split(bufio.ScanWords)
	if err != nil{
		log.Println("Error Encountered:", err)
	}
	defer f.Close()

	json := &jsonhandler.Metrics{}

	for i := 0; i < workers_count ; i++{
		wg.Go(func() {
			for job := range jobs{
				ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
				fetcher.Scraper(ctx, job, json)
				if ctx.Err() != nil{
					log.Println("Error Encountered:", ctx.Err())
					json.IncrementErrors()
				}
				cancel()
			}
		})
	}

	for br.Scan(){
		jobs <- br.Text()
	}

	close(jobs)
	wg.Wait()
	jsonhandler.Json_handle(json)
	go endpoint.Metric(json.Snapshot())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<- quit
	log.Println("Closing the endpoint!!")
	
}