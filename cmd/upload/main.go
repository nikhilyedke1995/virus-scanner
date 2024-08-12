package main

import (
	"log"

	"com.nikhil.virusscanner/internal/file/handler"
	"com.nikhil.virusscanner/internal/file/repository"
	"com.nikhil.virusscanner/lambda/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	cnfg, err := config.GetAWSConfig()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewS3Repository(cnfg.AwsSession)
	lambdaHandler := handler.NewHandler(*repo)
	lambda.Start(lambdaHandler.UploadObject)
}
