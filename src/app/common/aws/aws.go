package aws

import "github.com/aws/aws-sdk-go/aws"

type Aws struct {
	config *aws.Config
}

func NewAws(config *aws.Config) *Aws {
	return &Aws{
		config: config,
	}
}
