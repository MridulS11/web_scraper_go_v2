package fetcher

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"web_scraper_v2/configs"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(ctx context.Context, url string){
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil{
		log.Println("Error Encountered: ", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0 Safari/537.36")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil{
		log.Println("Error Encountered: ", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		log.Println("Error Encountered:", err)
	}

	base := filepath.Base(url)

	if strings.HasPrefix(base, "www."){
		base = strings.Replace(base, "www.", "", 1)
	}
	if filepath.Ext(base) != ""{
		base = strings.Replace(base, filepath.Ext(base), "", 1)
	}

	filename := base + ".txt"
		
	file, err := os.OpenFile(configs.OutPath+filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0622)
	if err != nil{
		log.Println("Error Encountered:", err)
	}
	br := bufio.NewWriter(file)

	doc.Find("#mw-content-text p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		if len(text) > 0{
			if _, err := br.WriteString(text); err != nil{
				log.Println("Error Encountered:", err)
				return
			}
			if _, err := br.WriteString("\n"); err != nil{
				log.Println("Error Encountered:", err)
				return
			}
		}
	})


}