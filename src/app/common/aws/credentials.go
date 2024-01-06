package aws

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func NewCredentials(options *CredentialsConfigOptions) *credentials.Credentials {
	if options.Static != nil {
		return newStaticCredentials(options.Static)
	}

	return nil
}

func newStaticCredentials(options *StaticCredentialsConfigOptions) *credentials.Credentials {
	return credentials.NewStaticCredentials(options.AccessKeyId, options.SecretAccessKey, options.Token)
}
