package domain

import (
	"boardType/internal/domain/vo"
	"errors"
	"time"
)

var _ BoardType = (*boardType)(nil)

type BoardType interface {
	ValidFiled() error
	ValidCreate() error
	ValidUpdate() error

	Update(name, description string) BoardType

	ToInfo() vo.Info
	ToDetail() vo.Detail
	ToUpdate() vo.Update
}

type boardType struct {
	id          int
	createBy    int
	cafeId      int
	name        string
	description string
	createdAt   time.Time
}

func (b *boardType) ToUpdate() vo.Update {
	return vo.Update{
		Id:          b.id,
		CreateBy:    b.createBy,
		CafeId:      b.cafeId,
		Name:        b.name,
		Description: b.description,
		CreatedAt:   b.createdAt,
	}
}

func (b *boardType) Update(name, description string) BoardType {
	b.name = name
	b.description = description
	return b
}

const (
	InvalidMemberId = "invalid member id"
	InvalidCafeId   = "invalid cafe id"
	InvalidName     = "invalid name"
	InvalidId       = "invalid id"
)

func (b *boardType) ValidCreate() error {
	if b.name == "" {
		return errors.New(InvalidName)
	}
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.createBy < 1 {
		return errors.New(InvalidMemberId)
	}
	return nil
}

func (b *boardType) ValidUpdate() error {
	if b.name == "" {
		return errors.New(InvalidName)
	}
	if b.id < 1 {
		return errors.New(InvalidId)
	}
	return nil
}

func (b *boardType) ValidFiled() error {
	if b.id < 1 {
		return errors.New(InvalidId)
	}
	if b.name == "" {
		return errors.New(InvalidName)
	}
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.createBy < 1 {
		return errors.New(InvalidMemberId)
	}
	return nil
}

func (b *boardType) ToInfo() vo.Info {
	return vo.Info{
		Id:          b.id,
		Name:        b.name,
		Description: b.description,
	}
}

func (b *boardType) ToDetail() vo.Detail {
	return vo.Detail{
		Id:          b.id,
		Name:        b.name,
		Description: b.description,
	}
}
