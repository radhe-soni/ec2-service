package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

//GetSession : returns an instance of aws session
func GetSession() *session.Session {
	ec2Region := viper.GetString("aws.ec2.region")
	accessKeyID := viper.GetString("aws.ec2.credentials.accessKeyId")
	accessKey := viper.GetString("aws.ec2.credentials.accessKey")
	creds := credentials.NewStaticCredentials(accessKeyID,
		accessKey, "")
	config := aws.Config{Region: aws.String(ec2Region), Credentials: creds}
	sess, err := session.NewSession(&config)
	if err != nil {
		log.Fatal("Error occured while getting session ", err)
	}
	log.Println("aws session obtained")
	return sess
}
