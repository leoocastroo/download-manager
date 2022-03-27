package utils

import (
	"net/http"
	"sync"
)


func downloadFile(link string, filename string, wg *sync.WaitGroup)  {
	resp, err := http.Get(link) // Get request

	if err != nil {
		print(err)
		return
	}

	defer resp.Body.Close() // Close request

	uploadToS3(resp.Body, filename)

	wg.Done() // decrement counter for goroutines
}

func DownloadFiles(filesMap map[string]string)  {
  // This WaitGroup is used to wait for all the goroutines launched here to finish
  var wg sync.WaitGroup
  for filename, link := range filesMap {
    wg.Add(1) // incrment counter for goroutines
    go downloadFile(link, filename, &wg)
  }
  wg.Wait()  // Main Goroutine will wait till incremnet counter is zero
}
