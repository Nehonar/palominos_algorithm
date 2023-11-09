package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/nehonar/palominos_algorithm/awsgo"
	"github.com/nehonar/palominos_algorithm/handlers"
	"github.com/nehonar/palominos_algorithm/models"
	secretmanager "github.com/nehonar/palominos_algorithm/secret_manager"
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

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error: wrong read secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return resp, nil
	}

	path := strings.Replace(request.PathParameters["palominos_algorithm"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.UserName)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	respApi := handlers.Handlers(awsgo.Ctx, request)

	if respApi.CustomResp != nil {
		return respApi.CustomResp, nil
	}

	resp = &events.APIGatewayProxyResponse{
		StatusCode: respApi.Status,
		Body:       respApi.Message,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
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
