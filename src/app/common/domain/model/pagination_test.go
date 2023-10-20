package model

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/exception"
)

func TestNewPageRequest(t *testing.T) {
	t.Run("returns PageRequest with values from query params", func(t *testing.T) {
		page := 2
		size := 5
		sortStrings := []string{"name,asc", "age,desc"}

		urlValues := url.Values{
			"page": {strconv.Itoa(page)},
			"size": {strconv.Itoa(size)},
			"sort": sortStrings,
		}

		expected := PageRequest{
			Page: page,
			Size: size,
			Sorts: []Sort{
				{Property: "name", Order: ORDER_ASC},
				{Property: "age", Order: ORDER_DESC},
			},
		}
		actual, err := NewPageRequest(urlValues)

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("returns PageRequest with default values when paging params are not informed", func(t *testing.T) {
		urlValues := url.Values{}

		expected := PageRequest{
			Size: DEFAULT_PAGE_SIZE,
		}

		actual, err := NewPageRequest(urlValues)

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
		_, err := NewPageRequest(urlValues)

		assert.Equal(t, expectedErr, err)
	})

	t.Run("throws illegal argument error when size param is invalid", func(t *testing.T) {
		param := "size"
		value := "s"

		urlValues := url.Values{
			param: {value},
		}

		expectedErr := exception.IllegalArgumentException(param, value)
		_, err := NewPageRequest(urlValues)

		assert.Equal(t, expectedErr, err)
	})

	t.Run("throws illegal argument error when sort param is invalid", func(t *testing.T) {
		param := "sort"
		value := "name,asc2"

		urlValues := url.Values{
			param: {value},
		}

		expectedErr := exception.IllegalArgumentException(param, value)
		_, err := NewPageRequest(urlValues)

		assert.Equal(t, expectedErr, err)
	})
}
