package urlhandlers

import (
	"fmt"
	"io"
	"log"
	"web_scraper_v2/configs"
	"web_scraper_v2/internals/fileHandlers"
	"web_scraper_v2/internals/worker"
)

func TerminalHandler(){
	var myArgs []string
	count := 1

	fmt.Println("Enter Stop() to start scraping.")


	for{
		var url string

		fmt.Printf("URL %d: ", count)

		_, err := fmt.Scan(&url)

		if url != "Exit()" || err == io.EOF{
			myArgs = append(myArgs, url)
		} else{
			break
		}

		if err != nil{
			log.Println(configs.ErrComment, err)
		}

		count++
	}
	fileHandlers.SliceToFile(myArgs)
	worker.ScraperPool()
}