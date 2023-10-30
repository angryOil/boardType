package model

import (
	"boardType/internal/domain"
	"github.com/uptrace/bun"
	"time"
)

type BoardType struct {
	bun.BaseModel `bun:"table:board_type,alias:bt"`

	Id          int       `bun:"id,pk,autoincrement"`
	CreateBy    int       `bun:"create_by,notnull"`
	CafeId      int       `bun:"cafe_id,notnull"`
	Name        string    `bun:"name,notnull"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at"`
}

func ToModel(d domain.BoardType) BoardType {
	return BoardType{
		Id:          d.Id,
		CreateBy:    d.CreateBy,
		CafeId:      d.CafeId,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
	}
}

func ToDomainLIst(models []BoardType) []domain.BoardType {
	results := make([]domain.BoardType, len(models))
	for i, m := range models {
		results[i] = m.ToDomain()
	}
	return results
}

func (m BoardType) ToDomain() domain.BoardType {
	return domain.BoardType{
		Id:          m.Id,
		CreateBy:    m.CreateBy,
		CafeId:      m.CafeId,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
	}
}
