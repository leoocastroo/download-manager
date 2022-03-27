package controller

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "download-manager/config"
  "download-manager/model"
  "download-manager/utils"
)


type CreateFileInput struct {
  Bucket  string `json:"bucket"`
  BucketKey string `json:"bucket_key"`
  Filename string `json:"filename" binding:"required"`
  URL string `json:"url" binding:"required"`
}


// GET /get-files-info
// Return all files saved
func GetFilesInfo(c *gin.Context)  {
  var files []model.File
  model.DB.Find(&files)

  c.JSON(http.StatusOK, gin.H{"data": files})
}

// POST /download-file
// Download a single file
func DownloadFile(c *gin.Context) {
  // Validate input
  var input CreateFileInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // save file
  file := model.File{
    Bucket:     config.GetEnvWithKey("BUCKET_NAME"),
    BucketKey:  config.GetEnvWithKey("BUCKET_KEY"),
    Filename:   input.Filename,
    URL:        input.URL,
  }

  // Save the infos about the file in the database
  model.DB.Create(&file)

  // Download and upload the file to s3
  utils.DownloadFile(file.URL, file.Filename)

  // Return the data
  c.JSON(http.StatusOK, gin.H{"data": file})
}

// POST /download-files
// Download multiple files
func DownloadFiles(c *gin.Context)  {
  var inputs[] CreateFileInput
  if err := c.ShouldBindJSON(&inputs); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  for i, _ := range inputs {
    inputs[i].Bucket = config.GetEnvWithKey("BUCKET_NAME")
    inputs[i].BucketKey = config.GetEnvWithKey("BUCKET_KEY")

    // Save the infos about the file in the database
    // model.DB.Create(&inputs[i])
  }

  // download and upload files to s3

  c.JSON(http.StatusOK, gin.H{"data": inputs})
}
