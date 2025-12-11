package urlhandlers

import (
	"fmt"
	"io"
	"log"
	"web_scraper_v2/configs"
)

func TerminalHandler(){
	var myArgs []string
	count := 1

	fmt.Println("Enter Exit() to exit the program")


	for{
		var url string

		fmt.Printf("URL %d: ", count)

		_, err := fmt.Scan(&url)

		if url != "Exit()"{
			myArgs = append(myArgs, url)
		}

		if err == io.EOF || url == "Exit()"{
			break
		}

		if err != nil{
			log.Println(configs.ErrComment, err)
		}

		count++
	}
	fmt.Println(myArgs)
}