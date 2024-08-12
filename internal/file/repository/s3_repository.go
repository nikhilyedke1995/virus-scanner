package repository

import (
	"bytes"
	"errors"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsS3Repository struct {
	S3Client s3.S3
}

func NewS3Repository(sess *session.Session) *AwsS3Repository {
	return &AwsS3Repository{
		S3Client: *s3.New(sess),
	}
}

func (repo *AwsS3Repository) PutObject(filename string, buffer []byte) (string, error) {
	_, err := repo.S3Client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:           aws.String(filename),
		Body:          bytes.NewReader(buffer),
		ContentType:   aws.String(http.DetectContentType(buffer)),
		ContentLength: aws.Int64(int64(len(buffer))),
	})
	if err != nil {
		return "", errors.New(err.Error())
	}
	return "object uploaded successfully", nil
}

func (repo *AwsS3Repository) GetObject(filename string) (*s3.GetObjectOutput, error) {
	getObject, err := repo.S3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(filename),
	})

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return getObject, nil
}
