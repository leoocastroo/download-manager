package main

import (
	"net/http"
	"sync"

	"download-manager/config"
	"download-manager/utils"
)


func downloadFile(link string, filename string, wg *sync.WaitGroup)  {
	resp, err := http.Get(link) // Get request

	if err != nil {
		print(err)
		return
	}

	defer resp.Body.Close() // Close request

	utils.UploadToS3(resp.Body, filename)

	wg.Done() // decrement counter for goroutines
}

func main() {
	config.LoadEnv() // load the enviroment

	m := map[string]string{
		"annual-enterprise-survey-2020-financial-year-provisional-csv.csv":"https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2020-financial-year-provisional/Download-data/annual-enterprise-survey-2020-financial-year-provisional-csv.csv"}

	// This WaitGroup is used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup
	for filename, link := range m {
		wg.Add(1) // incrment counter for goroutines
		go downloadFile(link, filename, &wg)
	}
	wg.Wait()  // Main Goroutine will wait till incremnet counter is zero
}
