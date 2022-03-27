package utils

import (
	"net/http"
	"sync"

)

// download multiple files at the same time and send to s3
func downloadFileAsync(link string, filename string, wg *sync.WaitGroup)  {
	resp, err := http.Get(link) // Get request

	if err != nil {
		print(err)
		return
	}

	defer resp.Body.Close() // Close request

	UploadToS3(resp.Body, filename)

	wg.Done() // decrement counter for goroutines
}

// download's control of multiple files
func DownloadFiles(filesMap map[string]string)  {
  // This WaitGroup is used to wait for all the goroutines launched here to finish
  var wg sync.WaitGroup
  for filename, link := range filesMap {
    wg.Add(1) // incrment counter for goroutines
    go downloadFileAsync(link, filename, &wg)
  }
  wg.Wait()  // Main Goroutine will wait till incremnet counter is zero
}

// download a single file and send to s3
func DownloadFile(link string, filename string)  {
  resp, err := http.Get(link) // Get request

  if err != nil {
    print(err)
    return
  }

  defer resp.Body.Close() // Close request

  UploadToS3(resp.Body, filename)
}
