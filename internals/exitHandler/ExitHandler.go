package exithandler

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
	"web_scraper_v2/configs"
)

func ExitFunc(){
	close := make(chan os.Signal, 1)
	signal.Notify(close, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<- close

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("\n" + "All the directories will be reset in 15 seconds to interupt press ctrl + c again")
	select{
	case <-time.After(15 * time.Second):
		fileSlice, err := filepath.Glob(filepath.Join(configs.UrlPath, "*.txt"))
		if err != nil{
			log.Println(configs.ErrComment, err)
		}
		for _, src := range fileSlice{
			if filepath.Base(src) == "cleaned_urls.txt"{
				continue
			}
			dst := filepath.Join(configs.RootPath, filepath.Base(src))
			
			if err := os.Rename(src, dst); err != nil{
				log.Println(configs.ErrComment, err)
			}
		}
		os.RemoveAll(configs.UrlPath)
		os.RemoveAll(configs.OutPath)
		log.Println("Bye!!")
	case <- quit:
		log.Println("Closing the endpoint!!")
	}
}