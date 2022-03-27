package model

type File struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Bucket string `json:"bucket"`
  BucketKey string `json:"bucket_key"`
  Filename  string `json:"filename"`
  URL string `json:"url"`
}
