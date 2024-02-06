package pagination

import (
	"fmt"

	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/errors"
)

func InvalidPageParam(param string, value interface{}) error {
	msg := fmt.Sprintf("Invalid value %q for param %s", value, param)
	return errors.BadRequestError(msg)
}
