package utils

import (
  "io"
  "log"
  "os"
  "download-manager/config"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
)


func connectAWS() *session.Session {
  var AccessKeyID string = config.GetEnvWithKey("AWS_ACCESS_KEY_ID")
  var SecretAccessKey string = config.GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
  var MyRegion string = config.GetEnvWithKey("AWS_REGION")

  log.Println("Region:", MyRegion)

  sess, err := session.NewSession(
    &aws.Config{
      Region: aws.String(MyRegion),
      Credentials: credentials.NewStaticCredentials(
        AccessKeyID,
        SecretAccessKey,
        "", // a token will be created when the session it's used.
      ),
    })

	if err != nil {
		panic(err)
	}
	return sess
}

func UploadToS3(file io.Reader, filename string)  {
  sess := connectAWS()
  uploader := s3manager.NewUploader(sess)

  var MyBucket string = config.GetEnvWithKey("BUCKET_NAME")
  var MyBucketKey string = config.GetEnvWithKey("BUCKET_KEY")

  log.Println("MyBucket:", MyBucket)
  log.Println("MyBucketKey:", MyBucketKey)

  _, err := uploader.Upload(&s3manager.UploadInput{
    Bucket: aws.String(MyBucket),
    Key:    aws.String(MyBucketKey+filename),
    Body:   file,
  })

  if err != nil {
    log.Fatalf("Error to upload file to s3: ", err)
    os.Exit(1)
  }
}
