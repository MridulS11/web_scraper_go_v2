package main

import (
	"fmt"
	"os"
	"web_scraper_v2/configs"
	"web_scraper_v2/internals/pathing"
)

func main(){
	err1, err2 := setup()
	if err1 != nil{
		fmt.Println("Error Encountered", err1)
	}
	if err2 != nil{
		fmt.Println("Error Encountered", err2)
	}
	fmt.Print("Enter 1 to work with a file with links(to be uploaded by you in the 'url' folder).\nEnter 2 to paste URLs in the terminal itself.\nInput: ")
	var choice int
	fmt.Scan(&choice)
	pathing.Path(choice)
}

func setup() (error, error){
	fmt.Println("Setting Up Project Directories...")
	return os.MkdirAll(configs.JsonPath, 0733), os.MkdirAll(configs.UrlPath, 0733)
}