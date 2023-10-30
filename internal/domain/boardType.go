package domain

import (
	"errors"
	"time"
)

type BoardType struct {
	Id          int
	CreateBy    int
	CafeId      int
	Name        string
	Description string
	CreatedAt   time.Time
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
