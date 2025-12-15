package fileHandlers

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"web_scraper_v2/configs"
)

func FileReader(s string){

	file, _ := os.Open(s)
	defer file.Close()

	r := bufio.NewScanner(file)
	r.Split(bufio.ScanWords)

	outPath := filepath.Join(configs.UrlPath, "cleaned_urls.txt")
	fileO, _ := os.OpenFile(outPath, os.O_CREATE | os.O_WRONLY, 0644)
	defer fileO.Close()
	w := bufio.NewWriter(fileO)

	fmt.Println("Cleaning The URLs...")
	for r.Scan(){
		w.WriteString(strings.Replace(r.Text(),",","",1)+"\n")
	}
	fmt.Println("URLs Cleaned, just in case you know")
	
	w.Flush()

}