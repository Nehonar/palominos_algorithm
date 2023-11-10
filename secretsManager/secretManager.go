package secretManager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/services/secretmanager"
	aws_go "github.com/nehonar/palominos_algorithm/awsGo"
	"github.com/nehonar/palominos_algorithm/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var secretData models.Secret
	fmt.Println("> secret " + secretName)

	svc := secretmanager.NewFromConfig(aws_go.Conf)
	key, err := svc.GetSecretValue(aws_go.Ctx, &secretmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)
	fmt.Println("> Read secret OK " + secretName)

	return secretData, nil
}
