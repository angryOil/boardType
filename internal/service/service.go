package service

import (
	"boardType/internal/domain"
	"boardType/internal/page"
	"boardType/internal/repository"
	"context"
	"errors"
)

type BoardTypeService struct {
	repo repository.BoardTypeRepository
}

func NewBoardTypeService(repo repository.BoardTypeRepository) BoardTypeService {
	return BoardTypeService{
		repo: repo,
	}
}
func (s BoardTypeService) Create(ctx context.Context, btd domain.BoardType) error {
	err := btd.ValidBoardType()
	if err != nil {
		return err
	}
	err = s.repo.Create(ctx, btd)
	return err
}

func (s BoardTypeService) GetListByCafe(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]domain.BoardType, int, error) {
	domains, total, err := s.repo.GetListByCafeId(ctx, cafeId, reqPage)
	return domains, total, err
}

func (s BoardTypeService) Delete(ctx context.Context, cafeId int, typeId int) error {
	err := s.repo.Delete(ctx, cafeId, typeId)
	return err
}

func (s BoardTypeService) Patch(ctx context.Context, d domain.BoardType) error {
	if d.Name == "" {
		return errors.New("invalid name")
	}
	err := s.repo.Patch(ctx, d.CafeId, d.Id,
		func(domains []domain.BoardType) (domain.BoardType, error) {
			if len(domains) == 0 {
				return domain.BoardType{}, errors.New("no rows")
			}
			return domains[0], nil
		},
		func(filtered domain.BoardType) domain.BoardType {
			return domain.BoardType{
				Id:          d.Id,
				CreateBy:    filtered.CreateBy,
				CafeId:      filtered.CafeId,
				Name:        d.Name,
				Description: d.Description,
				CreatedAt:   filtered.CreatedAt,
			}
		},
	)
	return err
}
