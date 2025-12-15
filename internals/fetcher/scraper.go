package fetcher

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"web_scraper_v2/configs"
	jsonhandler "web_scraper_v2/internals/jsonHandler"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(ctx context.Context, url string, json * jsonhandler.Metrics){
	
	start := time.Now()
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil{
		log.Println("Error Encountered: ", err)
		json.IncrementErrors()
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0 Safari/537.36")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil{
		log.Println("Error Encountered: ", err)
		json.IncrementErrors()
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		log.Println("Error Encountered:", err)
		json.IncrementErrors()
	}

	base := filepath.Base(url)

	if strings.HasPrefix(base, "https://"){
		base = strings.Replace(base, "www.", "", 1)
	}
	if strings.HasPrefix(base, "www."){
		base = strings.Replace(base, "www.", "", 1)
	}
	if filepath.Ext(base) != ""{
		base = strings.Replace(base, filepath.Ext(base), "", 1)
	}

	filename := base + ".txt"
		
	file, err := os.OpenFile(configs.OutPath+filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil{
		log.Println("Error Encountered:", err)
		json.IncrementErrors()
	} else{
		json.IncrementUrls()
	}

	defer file.Close()
	br := bufio.NewWriter(file)
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		if len(text) > 0{
			if _, err := br.WriteString(text); err != nil{
				log.Println("Error Encountered:", err)
				json.IncrementErrors()
				return
			} else {
				json.IncrementBytes(br.Size())
			}
			if _, err := br.WriteString("\n"); err != nil{
				log.Println("Error Encountered:", err)
				json.IncrementErrors()
				return
			}
		}
	})

	fmt.Println("Task Completed")
	json.IncrementTime(time.Since(start))
	br.Flush()
	//Execution End Of The Complete Scraper
}