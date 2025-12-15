package endpoint

import (
	"encoding/json"
	"log"
	"net/http"
	jsonhandler "web_scraper_v2/internals/jsonHandler"
)

func Metric(m jsonhandler.MetricsSnapshot){
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(m)
		if err != nil{
			panic(err)
		}
	})
	log.Println("Starting Server at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		panic(err)
	}
}