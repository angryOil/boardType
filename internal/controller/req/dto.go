package req

import (
	"boardType/internal/domain"
	"time"
)

type CreateBoardTypeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d CreateBoardTypeDto) ToDomain(cafeId, memberId int) domain.BoardType {
	return domain.BoardType{
		CreateBy:    memberId,
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   time.Now(),
	}
}

type PatchBoardDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d PatchBoardDto) ToDomain(cafeId, typeId int) domain.BoardType {
	return domain.BoardType{
		Id:          typeId,
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
	}
}
