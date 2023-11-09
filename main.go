package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/nehonar/palominos_algorithm/awsgo"
)

func main() {
	lambda.Start(lambdaStart)
}

func lambdaStart(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var resp *events.APIGatewayProxyResponse

	awsgo.AWSInit()

	if !validateParams() {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid params. You need: secretName, bucketName, urlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}

}

func validateParams() bool {
	_, checkParam := os.LookupEnv("SecretName")
	if !checkParam {
		return checkParam
	}

	_, checkParam = os.LookupEnv("BucketName")
	if !checkParam {
		return checkParam
	}

	_, checkParam = os.LookupEnv("UrlPrefix")
	if !checkParam {
		return checkParam
	}

	return checkParam
}
