package domain

import "time"

var _ BoardTypeBuilder = (*boardTypeBuilder)(nil)

func NewBoardTypeBuilder() BoardTypeBuilder {
	return &boardTypeBuilder{}
}

type BoardTypeBuilder interface {
	Id(id int) BoardTypeBuilder
	CreateBy(createdBy int) BoardTypeBuilder
	CafeId(cafeId int) BoardTypeBuilder
	Name(name string) BoardTypeBuilder
	Description(description string) BoardTypeBuilder
	CreatedAt(createdAt time.Time) BoardTypeBuilder

	Build() BoardType
}

type boardTypeBuilder struct {
	id          int
	createBy    int
	cafeId      int
	name        string
	description string
	createdAt   time.Time
}

func (b *boardTypeBuilder) Id(id int) BoardTypeBuilder {
	b.id = id
	return b
}

func (b *boardTypeBuilder) CreateBy(createdBy int) BoardTypeBuilder {
	b.createBy = createdBy
	return b
}

func (b *boardTypeBuilder) CafeId(cafeId int) BoardTypeBuilder {
	b.cafeId = cafeId
	return b
}

func (b *boardTypeBuilder) Name(name string) BoardTypeBuilder {
	b.name = name
	return b
}

func (b *boardTypeBuilder) Description(description string) BoardTypeBuilder {
	b.description = description
	return b
}

func (b *boardTypeBuilder) CreatedAt(createdAt time.Time) BoardTypeBuilder {
	b.createdAt = createdAt
	return b
}

func (b *boardTypeBuilder) Build() BoardType {
	return &boardType{
		id:          b.id,
		createBy:    b.createBy,
		cafeId:      b.cafeId,
		name:        b.name,
		description: b.description,
		createdAt:   b.createdAt,
	}
}
