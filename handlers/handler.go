package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
