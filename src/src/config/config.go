package config

import (
  "log"
  "os"

  "github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
  return os.Getenv(key) // get env system
}

func LoadEnv()  {
  err := godotenv.Load("../config/.env") // load .env file into memory
  if err != nil {
    log.Fatalf("Error loading .env file")
    os.Exit(1)
  }
}
