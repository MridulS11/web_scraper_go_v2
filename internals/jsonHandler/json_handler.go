package jsonhandler

import (
	"bufio"
	"encoding/json"
	"os"
	"web_scraper_v2/configs"
)


func Json_handle(m * Metrics){

	snap := m.Snapshot()
	res2, _ := json.MarshalIndent(snap, "", " ")

	file, _ := os.OpenFile(configs.JsonPath + "metrics.json", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	br := bufio.NewWriter(file)
	br.Write(res2)
	defer file.Close()
	br.Flush()

}