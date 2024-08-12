package handler

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"com.nikhil.virusscanner/internal/file/repository"
	"github.com/aws/aws-lambda-go/events"
	"github.com/grokify/go-awslambda"
)

type Handler struct {
	awsRepo repository.AwsS3Repository
}

func NewHandler(awsRepo repository.AwsS3Repository) *Handler {
	return &Handler{awsRepo: awsRepo}
}

func (handler *Handler) UploadObject(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{}
	r, err := awslambda.NewReaderMultipart(request)
	if err != nil {
		return res, err
	}
	part, err := r.NextPart()
	if err != nil {
		return res, err
	}
	content, err := io.ReadAll(part)
	if err != nil {
		return res, err
	}
	obj, err := handler.awsRepo.PutObject(part.FileName(), content)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	response := make(map[string]string)
	response["message"] = obj

	marshal, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	res = events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(marshal),
	}
	return res, nil
}

func (handler *Handler) View(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{}
	key := request.QueryStringParameters["filename"]

	getObject, err := handler.awsRepo.GetObject(key)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	defer getObject.Body.Close()
	content, err := ioutil.ReadAll(getObject.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	base64Content := base64.StdEncoding.EncodeToString(content)
	res = events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "image/png",
		},
		Body:            base64Content,
		IsBase64Encoded: true,
	}
	return res, nil
}
