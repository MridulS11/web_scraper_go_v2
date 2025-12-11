package urlhandlers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func UploadHandler(){

	file, _ := os.Getwd()
	fullpath := filepath.Join(file, "url")

	fmt.Println("Looking For File Upload...")

	var matches []string

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for{
		select{
		case <- timeout:
			fmt.Println("File Not Received")
			os.Exit(1)
		case <- ticker.C:
			matches, _ = filepath.Glob(filepath.Join(fullpath, "*.txt"))
			if len(matches) > 0 {
				fmt.Println("File Received:", matches[0])
				return
			}
		}
	}
	
}