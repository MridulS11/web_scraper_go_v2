package main

import (
	"fmt"
	"web_scraper_v2/internals/pathing"
)

func main(){
	fmt.Print("Enter 1 to work with a file with links(to be uploaded by you in the links folder).\nEnter 2 to paste URLs in the terminal itself.\nInput: ")
	var choice int
	fmt.Scan(&choice)
	pathing.Path(choice)
}