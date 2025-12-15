package pathing

import (
	"os"
	urlhandlers "web_scraper_v2/internals/urlHandlers"
)

func Path(choice int){
	
	switch choice{
	case 1:
		urlhandlers.UploadHandler()
	case 2:
		urlhandlers.TerminalHandler()
	default:
		os.Exit(1)
	}
}