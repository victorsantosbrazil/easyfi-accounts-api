package pagination

import (
	"strconv"

	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/exception"
)

type QueryParams map[string][]string

func (p QueryParams) GetString(key string) string {
	if len(p[key]) == 0 {
		return ""
	}

	return p[key][0]
}

func (p QueryParams) GetStrings(key string) []string {
	return p[key]
}

func (p QueryParams) GetIntOrDefault(key string, dflt int) (value int, err error) {
	if strValue := p.GetString(key); strValue != "" {
		value, err = strconv.Atoi(strValue)
		if err != nil {
			return 0, exception.IllegalArgumentException(key, value)
		}
		return value, nil
	} else {
		return dflt, nil
	}
}
