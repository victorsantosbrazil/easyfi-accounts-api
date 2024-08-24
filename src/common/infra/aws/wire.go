package aws

import (
	"github.com/google/wire"
)

var AwsSet = wire.NewSet(NewAwsConfig, NewAws)
