package exception

import (
	"fmt"
)

func IllegalArgumentException(argument string, value interface{}) error {
	return fmt.Errorf("invalid value %q for argument %s", value, argument)
}
