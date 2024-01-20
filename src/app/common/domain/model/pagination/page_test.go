package pagination

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageMap(t *testing.T) {
	pagination := Pagination{Page: 1, Size: 10}
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	intsPage := Page[int]{Pagination: pagination, Items: items}
	mapFn := func(i int) string { return strconv.Itoa(i) }

	strsPage := MapPage(intsPage, mapFn)

	assert.Equal(t, pagination, strsPage.Pagination)
	for i := 0; i < len(items); i++ {
		assert.Equal(t, mapFn(items[i]), strsPage.Items[i])
	}
}
