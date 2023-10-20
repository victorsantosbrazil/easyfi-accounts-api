package model

import (
	"net/url"
	"strings"

	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/exception"
)

const (
	ORDER_ASC  = "ASC"
	ORDER_DESC = "DESC"
)

const (
	DEFAULT_PAGE_SIZE = 10
)

type (
	PageRequest struct {
		Page  int
		Size  int
		Sorts []Sort
	}

	Sort struct {
		Property string
		Order    string
	}

	Pagination struct {
		Page       int  `json:"page"`
		Size       int  `json:"size"`
		TotalPages int  `json:"totalPages"`
		Total      int  `json:"total"`
		Last       bool `json:"last"`
		First      bool `json:"first"`
	}
)

func NewPageRequest(urlValues url.Values) (pageRequest PageRequest, err error) {
	queryParams := QueryParams(urlValues)

	pageRequest.Page, err = queryParams.GetIntOrDefault("page", 0)
	if err != nil {
		return pageRequest, err
	}

	pageRequest.Size, err = queryParams.GetIntOrDefault("size", DEFAULT_PAGE_SIZE)
	if err != nil {
		return pageRequest, err
	}

	pageRequest.Sorts, err = newPageRequestSorts(queryParams)

	return pageRequest, err
}

func newPageRequestSorts(queryParams QueryParams) (sorts []Sort, err error) {
	sortStrings := queryParams.GetStrings("sort")

	for _, sortString := range sortStrings {
		sort, err := newSort(sortString)
		if err != nil {
			return sorts, err
		}
		sorts = append(sorts, sort)
	}

	return sorts, nil
}

func newSort(sortString string) (Sort, error) {
	parts := strings.Split(sortString, ",")
	property := parts[0]

	var order string
	if len(parts) > 1 {
		order = strings.ToUpper(parts[1])
	} else {
		order = ORDER_ASC
	}

	if order != ORDER_ASC && order != ORDER_DESC {
		return Sort{}, exception.IllegalArgumentException("sort", sortString)
	}

	sort := Sort{
		Property: property,
		Order:    order,
	}
	return sort, nil
}
