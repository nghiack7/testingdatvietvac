package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Config struct {
	Address string
	Region  string
	Profile string
	ID      string
	Secret  string
}

func ConnectAWS(config Config) *session.Session {

	sess, err := session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Region:           aws.String(config.Region),
				Credentials:      credentials.NewStaticCredentials(config.ID, config.Secret, ""),
				Endpoint:         aws.String(config.Address),
				S3ForcePathStyle: aws.Bool(true),
			},
			Profile: config.Profile,
		})
	if err != nil {
		panic(err)
	}
	return sess
}
