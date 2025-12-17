package exithandler

import (
	"syscall"
	"os/signal"
	"fmt"
	"web_scraper_v2/configs"
	"os"
	"time"
	"log"
)

func ExitFunc(){
	close := make(chan os.Signal, 1)
	signal.Notify(close, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<- close

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("\n" + "All the directories will be reset in 30 seconds or press ctrl + c again")
	select{
	case <-time.After(30 * time.Second):
		os.RemoveAll(configs.UrlPath)
		os.RemoveAll(configs.OutPath)
	case <- quit:
		log.Println("Closing the endpoint!!")
	}
}