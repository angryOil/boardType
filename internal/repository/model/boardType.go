package model

import (
	"boardType/internal/domain"
	"boardType/internal/repository/req"
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

func ToCreateModel(c req.Create) BoardType {
	return BoardType{
		CreateBy:    c.CreateBy,
		CafeId:      c.CafeId,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
	}
}

func ToUpdateModel(u req.Update) BoardType {
	return BoardType{
		Id:          u.Id,
		CreateBy:    u.MemberId,
		CafeId:      u.CafeId,
		Name:        u.Name,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
	}
}

func ToDomainList(models []BoardType) []domain.BoardType {
	results := make([]domain.BoardType, len(models))
	for i, m := range models {
		results[i] = m.ToDomain()
	}
	return results
}

func (m BoardType) ToDomain() domain.BoardType {
	return domain.NewBoardTypeBuilder().
		Id(m.Id).
		CreateBy(m.CreateBy).
		CafeId(m.CafeId).
		Name(m.Name).
		Description(m.Description).
		CreatedAt(m.CreatedAt).
		Build()
}
