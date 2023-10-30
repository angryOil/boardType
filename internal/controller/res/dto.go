package res

import "boardType/internal/domain"

type BoardTypeDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToBoardTypeDtoList(domains []domain.BoardType) []BoardTypeDto {
	results := make([]BoardTypeDto, len(domains))
	for i, d := range domains {
		results[i] = BoardTypeDto{
			Id:          d.Id,
			Name:        d.Name,
			Description: d.Description,
		}
	}
	return results
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
