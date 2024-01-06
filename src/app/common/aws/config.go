package aws

import (
	"github.com/aws/aws-sdk-go/aws"
)

type ConfigOptions struct {
	Region      string
	Endpoint    string
	Credentials *CredentialsConfigOptions
}

type CredentialsConfigOptions struct {
	Static *StaticCredentialsConfigOptions
}

type StaticCredentialsConfigOptions struct {
	AccessKeyId     string
	SecretAccessKey string
	Token           string
}

func NewAwsConfig(options *ConfigOptions) *aws.Config {
	return &aws.Config{
		Region:      aws.String(options.Region),
		Endpoint:    aws.String(options.Endpoint),
		Credentials: NewCredentials(options.Credentials),
	}
}
