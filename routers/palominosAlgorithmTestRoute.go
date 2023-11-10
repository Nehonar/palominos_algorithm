package routers

import (
	"context"
	"io"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nehonar/palominos_algorithm/models"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func PalominosAlgotrithmTest(ctx context.Context, request events.APIGatewayProxyRequest) models.ApiResponse {

	return models.ApiResponse{
		Status:  200,
		Message: "Archivo recibido",
	}
}
