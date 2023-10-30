package domain

import (
	"errors"
	"time"
)

type BoardType struct {
	Id          int       `json:"id"`
	CreateBy    int       `json:"create_by"`
	CafeId      int       `json:"cafe_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (bt BoardType) ValidBoardType() error {
	if bt.CreateBy == 0 {
		return errors.New("invalid member id")
	}
	if bt.CafeId == 0 {
		return errors.New("invalid cafe id")
	}
	if bt.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}
