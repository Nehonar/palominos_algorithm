package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Verifica si la ruta es /ping
	if request.HTTPMethod == "GET" && request.Path == "/ping" {
		resp := Response{Message: "pong"}
		body, _ := json.Marshal(resp)

		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(body),
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       `{"message": "Not found"}`,
	}, nil
}

func main() {
	lambda.Start(handler)
}
