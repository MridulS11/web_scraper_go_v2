package urlhandlers

import (
	"fmt"
	"io"
	"log"
	"web_scraper_v2/configs"
	fileslice "web_scraper_v2/internals/fileSlice"
	"web_scraper_v2/internals/worker"
)

func TerminalHandler(){
	var myArgs []string
	count := 1

	fmt.Println("Enter Exit() to exit the program")


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
	fileslice.SliceToFile(myArgs)
	worker.ScraperPool()
}