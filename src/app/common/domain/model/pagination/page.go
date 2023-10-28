package pagination

type (
	Pagination struct {
		Page       int  `json:"page"`
		Size       int  `json:"size"`
		TotalPages int  `json:"totalPages"`
		Total      int  `json:"total"`
		Last       bool `json:"last"`
		First      bool `json:"first"`
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
