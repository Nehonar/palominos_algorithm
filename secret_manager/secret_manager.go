package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/nehonar/palominos_algorithm/awsgo"
	"github.com/nehonar/palominos_algorithm/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var secret_data models.Secret
	fmt.Println("> secret " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Conf)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secret_data, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secret_data)
	fmt.Println("> Read secret OK " + secretName)

	return secret_data, nil
}
