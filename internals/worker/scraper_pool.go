package worker

import (
	"bufio"
	"context"
	"log"
	"os"
	"sync"
	"time"
	"web_scraper_v2/configs"
	"web_scraper_v2/internals/fetcher"
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

	for i := 0; i < workers_count ; i++{
		wg.Go(func() {
			for job := range jobs{
				ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
				fetcher.Scraper(ctx, job)
				if ctx.Err() != nil{
					log.Println("Error Encountered:", ctx.Err())
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
}