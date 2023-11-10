package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	awsGo "github.com/nehonar/palominos_algorithm/awsGo"
	"github.com/nehonar/palominos_algorithm/handlers"
	"github.com/nehonar/palominos_algorithm/models"
	secretsManager "github.com/nehonar/palominos_algorithm/secretsManager"
)

func main() {
	lambda.Start(lambdaStart)
}

func lambdaStart(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var resp *events.APIGatewayProxyResponse

	awsGo.AWSInit()

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

	SecretModel, err := secretsManager.GetSecret(os.Getenv("SecretName"))
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

	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("path"), path)
	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("method"), request.HTTPMethod)
	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("user"), SecretModel.UserName)
	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("password"), SecretModel.Password)
	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("host"), SecretModel.Host)
	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("database"), SecretModel.Database)
	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("body"), request.Body)
	awsGo.Ctx = context.WithValue(awsGo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	respApi := handlers.Handlers(awsGo.Ctx, request)

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
