package pagination

import (
	"fmt"

	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/errors"
)

func InvalidPageParam(param string, value interface{}) error {
	msg := fmt.Sprintf("Invalid value %q for %s param", value, param)
	return errors.BadRequestError(msg)
}
