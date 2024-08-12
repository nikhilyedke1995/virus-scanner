package config

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

// aws properties
type AwsProps struct {
	AWS_REGION      string
	AWS_BUCKET_NAME string
}

// aws config contains session and properties
type AwsConfig struct {
	AwsSession *session.Session
	AwsProps   *AwsProps
}

func GetAWSConfig() (*AwsConfig, error) {
	//loading config
	loadConfig()
	//creating session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS"),
			os.Getenv("AWS_SECRET"),
			"",
		),
	})
	return &AwsConfig{
		AwsSession: sess,
		AwsProps: &AwsProps{
			AWS_REGION:      os.Getenv("REGION"),
			AWS_BUCKET_NAME: os.Getenv("AWS_BUCKET_NAME"),
		},
	}, err
}

func loadConfig() {
	err := godotenv.Load("./lambda/config/.env")
	if err != nil {
		//log.Fatal("error loading env file")
		fmt.Println("error loading env file")
	}
}
