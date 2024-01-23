package pagination

type (
	Pagination struct {
		Page          int `json:"page" example:"1"`
		Size          int `json:"size" example:"10"`
		TotalPages    int `json:"totalPages" example:"5"`
		TotalElements int `json:"totalElements" example:"50"`
	}

	Page[T any] struct {
		Pagination Pagination `json:"pagination"`
		Items      []T        `json:"items"`
	}
)

func MapPage[M, N any](src Page[M], mapFn func(M) N) Page[N] {
	items := make([]N, len(src.Items))

	for i, item := range src.Items {
		items[i] = mapFn(item)
	}

	return Page[N]{
		Pagination: src.Pagination,
		Items:      items,
	}
}
