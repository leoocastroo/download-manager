package main

import (
	"download-manager/config"
	"download-manager/utils"
)




func main() {
	config.LoadEnv() // load the enviroment

	m := map[string]string{
   "business-financial-data-december-2021-quarter-csv.zip": "https://www.stats.govt.nz/assets/Uploads/Business-financial-data/Business-financial-data-December-2021-quarter/Download-data/business-financial-data-december-2021-quarter-csv.zip"}

	utils.DownloadFiles(m)
}
