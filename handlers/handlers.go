package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nehonar/palominos_algorithm/jwt"
	"github.com/nehonar/palominos_algorithm/models"
	"github.com/nehonar/palominos_algorithm/public"
	"github.com/nehonar/palominos_algorithm/routers"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ApiResponse {
	fmt.Println("Processing" + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var resp models.ApiResponse
	resp.Status = 400

	isOK, statusCode, msg, _ := authorization(ctx, request)
	if !isOK {
		resp.Status = statusCode
		resp.Message = msg
		return resp
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "/index":
			htmlContent, err := public.LoadIndexFile("palominos_algorithm/public/index.html")
			if err != nil {
				resp.Status = 500
				resp.Message = "Invalid url "
				return resp
			}

			resp.Status = 200
			resp.CustomResp.Body = htmlContent
			resp.CustomResp.Headers = map[string]string{
				"Content-Type": "text/html",
			}
			return resp
		case "palominosAlgotrithmTest":
			return routers.PalominosAlgotrithmTest(ctx, request)
		}
	}

	resp.Message = "Invalid method"

	return resp
}

func authorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "palominosAlgotrithmTest" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Required token", models.Claim{}
	}

	claim, ok, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("JWTSign")).(string))
	if !ok {
		if err != nil {
			fmt.Println("Token error " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Token error " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")

	return true, 200, msg, *claim
}
