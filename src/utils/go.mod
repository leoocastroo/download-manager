module download-manager/utils

go 1.18

require (
	download-manager/config v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go v1.43.25
)

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
)

replace download-manager/config => ../config
