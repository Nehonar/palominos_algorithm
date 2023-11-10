package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nehonar/palominos_algorithm/models"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ApiResponse {
	fmt.Println("Processing" + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var resp models.ApiResponse
	resp.Status = 400

	isOK, statusCode, msg, claim := authorization(ctx, request)

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	resp.Message = "Invalid method"

	return resp
}

func authorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "uploadFile" {
		return true, 200, "", models.Claim{}
	}
}
