package utils

import (
  "download-manager/config"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
)


func connectAWS() *session.Session {
  var AccessKeyID string = GetEnvWithKey("AWS_ACCESS_KEY_ID")
  var SecretAccessKey string = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
  var MyRegion string = GetEnvWithKey("AWS_REGION")

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

func UploadToS3(file)  {
  var s3 *session.Session = connectAWS()
  
}
