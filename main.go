package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nehonar/palominos_algorithm/handlers"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(handlers.Handler())
}
