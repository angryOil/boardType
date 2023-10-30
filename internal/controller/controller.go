package controller

import (
	"boardType/internal/controller/req"
	"boardType/internal/controller/res"
	"boardType/internal/page"
	"boardType/internal/service"
	"context"
)

type BoardTypeController struct {
	s service.BoardTypeService
}

func NewBoardTypeController(s service.BoardTypeService) BoardTypeController {
	return BoardTypeController{
		s: s,
	}
}
func (c BoardTypeController) Create(ctx context.Context, cafeId int, memberId int, d req.CreateBoardTypeDto) error {
	btd := d.ToDomain(cafeId, memberId)
	err := c.s.Create(ctx, btd)
	return err
}

func (c BoardTypeController) GetListByCafe(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.BoardTypeDto, int, error) {
	domains, total, err := c.s.GetListByCafe(ctx, cafeId, reqPage)
	if err != nil {
		return []res.BoardTypeDto{}, 0, err
	}
	return res.ToBoardTypeDtoList(domains), total, nil
}
