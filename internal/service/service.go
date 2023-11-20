package service

import (
	"boardType/internal/domain"
	"boardType/internal/domain/vo"
	"boardType/internal/page"
	"boardType/internal/repository"
	req2 "boardType/internal/repository/req"
	"boardType/internal/service/req"
	"boardType/internal/service/res"
	"context"
	"errors"
	"time"
)

type BoardTypeService struct {
	repo repository.BoardTypeRepository
}

func NewBoardTypeService(repo repository.BoardTypeRepository) BoardTypeService {
	return BoardTypeService{
		repo: repo,
	}
}

const (
	NoRows    = "no rows"
	InvalidId = "invalid board type id"
)

func (s BoardTypeService) Create(ctx context.Context, c req.Create) error {
	name, description := c.Name, c.Description
	cafeId, memberId := c.CafeId, c.MemberId
	createdAt := time.Now()
	err := domain.NewBoardTypeBuilder().
		Name(name).
		Description(description).
		CafeId(cafeId).
		CreateBy(memberId).
		CreatedAt(createdAt).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, req2.Create{
		Name:        name,
		Description: description,
		CafeId:      cafeId,
		CreateBy:    memberId,
		CreatedAt:   createdAt,
	})
	return err
}

func (s BoardTypeService) GetListByCafe(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.GetListByCafe, int, error) {
	domains, total, err := s.repo.GetListByCafeId(ctx, cafeId, reqPage)
	dto := make([]res.GetListByCafe, len(domains))
	for i, d := range domains {
		v := d.ToInfo()
		dto[i] = res.GetListByCafe{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
		}
	}
	return dto, total, err
}

func (s BoardTypeService) Delete(ctx context.Context, cafeId int, typeId int) error {
	err := s.repo.Delete(ctx, cafeId, typeId)
	return err
}

func (s BoardTypeService) Patch(ctx context.Context, p req.Patch) error {
	id, cafeId, createBy := p.Id, p.CafeId, p.MemberId
	name, description := p.Name, p.Description
	err := domain.NewBoardTypeBuilder().
		Id(id).
		Name(name).
		Description(description).
		CafeId(cafeId).
		CreateBy(createBy).
		Build().ValidUpdate()

	err = s.repo.Patch(ctx, p.Id,
		func(domains []domain.BoardType) (domain.BoardType, error) {
			if len(domains) != 1 {
				return nil, errors.New(NoRows)
			}
			return domains[0], nil
		},
		func(filtered domain.BoardType) vo.Update {
			u := filtered.Update(name, description)
			return u.ToUpdate()
		},
	)
	return err
}

func (s BoardTypeService) GetDetail(ctx context.Context, id int) (res.GetDetail, error) {
	if id < 1 {
		return res.GetDetail{}, errors.New(InvalidId)
	}
	d, err := s.repo.GetDetail(ctx, id)
	if err != nil {
		return res.GetDetail{}, err
	}
	v := d.ToDetail()
	return res.GetDetail{
		Id:          v.Id,
		Name:        v.Name,
		Description: v.Description,
	}, nil
}
