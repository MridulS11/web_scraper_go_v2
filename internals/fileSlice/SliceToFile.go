package fileslice

import (
	"bufio"
	"os"
	"path/filepath"
	"web_scraper_v2/configs"
)

func SliceToFile(s []string){

	outPath := filepath.Join(configs.UrlPath, "cleaned_urls.txt")
	fileO, _ := os.OpenFile(outPath, os.O_CREATE | os.O_WRONLY, 0644)
	defer fileO.Close()
	w := bufio.NewWriter(fileO)

	for _, url := range s{
		w.WriteString(url+"\n")
	}

	w.Flush()

}