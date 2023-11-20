package controller

import (
	"boardType/internal/controller/req"
	"boardType/internal/controller/res"
	"boardType/internal/page"
	"boardType/internal/service"
	req2 "boardType/internal/service/req"
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
	err := c.s.Create(ctx, req2.Create{
		Name:        d.Name,
		Description: d.Description,
		CafeId:      cafeId,
		MemberId:    memberId,
	})
	return err
}

func (c BoardTypeController) GetListByCafe(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.BoardTypeDto, int, error) {
	listArr, total, err := c.s.GetListByCafe(ctx, cafeId, reqPage)
	if err != nil {
		return []res.BoardTypeDto{}, 0, err
	}
	dto := make([]res.BoardTypeDto, len(listArr))
	for i, l := range listArr {
		dto[i] = res.BoardTypeDto{
			Id:          l.Id,
			Name:        l.Name,
			Description: l.Description,
		}
	}
	return dto, total, nil
}

func (c BoardTypeController) Delete(ctx context.Context, cafeId int, typeId int) error {
	err := c.s.Delete(ctx, cafeId, typeId)
	return err
}

func (c BoardTypeController) Patch(ctx context.Context, cafeId, typeId int, d req.PatchBoardDto) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:          typeId,
		Name:        d.Name,
		Description: d.Description,
		CafeId:      cafeId,
	})
	return err
}

func (c BoardTypeController) GetDetail(ctx context.Context, id int) (res.BoardDetailDto, error) {
	d, err := c.s.GetDetail(ctx, id)
	if err != nil {
		return res.BoardDetailDto{}, err
	}
	return res.BoardDetailDto{
		Id:          d.Id,
		Name:        d.Name,
		Description: d.Description,
	}, nil
}
