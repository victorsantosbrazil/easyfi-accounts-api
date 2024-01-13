package pagination

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/exception"
)

func TestNewPageParams(t *testing.T) {
	t.Run("returns PageParams with values from query params", func(t *testing.T) {
		page := 2
		size := 5
		sortStrings := []string{"name,asc", "age,desc"}

		urlValues := url.Values{
			"page": {strconv.Itoa(page)},
			"size": {strconv.Itoa(size)},
			"sort": sortStrings,
		}

		expected := PageParams{
			Page: page,
			Size: size,
			Sorts: []Sort{
				{Property: "name", Order: ORDER_ASC},
				{Property: "age", Order: ORDER_DESC},
			},
		}
		actual, err := NewPageParams(urlValues)

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("returns PageParams with default values when paging params are not informed", func(t *testing.T) {
		urlValues := url.Values{}

		expected := PageParams{
			Page: 1,
			Size: DEFAULT_PAGE_SIZE,
		}

		actual, err := NewPageParams(urlValues)

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("throws illegal argument error when page param is invalid", func(t *testing.T) {
		param := "page"
		value := "s"

		urlValues := url.Values{
			param: {value},
		}

		expectedErr := exception.IllegalArgumentException(param, value)
		_, err := NewPageParams(urlValues)

		assert.Equal(t, expectedErr, err)
	})

	t.Run("throws illegal argument error when size param is invalid", func(t *testing.T) {
		param := "size"
		value := "s"

		urlValues := url.Values{
			param: {value},
		}

		expectedErr := exception.IllegalArgumentException(param, value)
		_, err := NewPageParams(urlValues)

		assert.Equal(t, expectedErr, err)
	})

	t.Run("throws illegal argument error when sort param is invalid", func(t *testing.T) {
		param := "sort"
		value := "name,asc2"

		urlValues := url.Values{
			param: {value},
		}

		expectedErr := exception.IllegalArgumentException(param, value)
		_, err := NewPageParams(urlValues)

		assert.Equal(t, expectedErr, err)
	})
}
