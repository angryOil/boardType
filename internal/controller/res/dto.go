package res

type BoardTypeDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ListTotalDto[T any] struct {
	Contents []T `json:"contents,omitempty"`
	Total    int `json:"total,omitempty"`
}

func NewListTotalDto[T any](contents []T, total int) ListTotalDto[T] {
	return ListTotalDto[T]{
		Contents: contents,
		Total:    total,
	}
}
