package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Conf aws.Config
var err error

func AWSInit() {
	Ctx = context.TODO()
	Conf, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("eu-north-1"))
	if err != nil {
		panic("Error to download config .aws/config " + err.Error())
	}
}
