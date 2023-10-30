package service

import (
	"boardType/internal/domain"
	"boardType/internal/page"
	"boardType/internal/repository"
	"context"
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
