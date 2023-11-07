package repository

import (
	"boardType/internal/domain"
	"boardType/internal/domain/vo"
	"boardType/internal/page"
	"boardType/internal/repository/model"
	req2 "boardType/internal/repository/req"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
)

type BoardTypeRepository struct {
	db bun.IDB
}

func NewBoardTypeRepository(db bun.IDB) BoardTypeRepository {
	return BoardTypeRepository{
		db: db,
	}
}

const (
	InternalServerError = "internal server error"
)

func (r BoardTypeRepository) Create(ctx context.Context, c req2.Create) error {
	m := model.ToCreateModel(c)
	_, err := r.db.NewInsert().Model(&m).Exec(ctx)
	return err
}

func (r BoardTypeRepository) GetListByCafeId(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]domain.BoardType, int, error) {
	var models []model.BoardType
	cnt, err := r.db.NewSelect().Model(&models).
		ColumnExpr("name,id,substring(bt.description,1,50) as description").
		Where("cafe_id = ?", cafeId).Limit(reqPage.Size).Offset(reqPage.OffSet).Order("id desc").
		ScanAndCount(ctx)
	if err != nil {
		log.Println("GetListByCafeId NewSelect err: ", err)
		return []domain.BoardType{}, 0, errors.New(InternalServerError)
	}
	return model.ToDomainList(models), cnt, nil
}

func (r BoardTypeRepository) Delete(ctx context.Context, cafeId int, typeId int) error {
	var m model.BoardType
	_, err := r.db.NewDelete().Model(&m).Where("id =? and cafe_id = ?", typeId, cafeId).Exec(ctx)
	if err != nil {
		log.Println("Delete NewDelete err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}

func (r BoardTypeRepository) Patch(
	ctx context.Context, id int,
	validFunc func(domains []domain.BoardType) (domain.BoardType, error),
	mergeFunc func(filtered domain.BoardType) vo.Update) error {
	var models []model.BoardType
	err := r.db.NewSelect().Model(&models).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("Patch NewSelect err: ", err)
		return errors.New(InternalServerError)
	}
	validD, err := validFunc(model.ToDomainList(models))
	if err != nil {
		return err
	}

	updatedVo := mergeFunc(validD)

	mergedModel := model.ToUpdateModel(req2.Update{
		Id:          updatedVo.Id,
		Name:        updatedVo.Name,
		Description: updatedVo.Description,
		CafeId:      updatedVo.CafeId,
		MemberId:    updatedVo.CreateBy,
		CreatedAt:   updatedVo.CreatedAt,
	})

	_, err = r.db.NewInsert().Model(&mergedModel).On("CONFLICT (id) DO UPDATE").Exec(ctx)
	if err != nil {
		log.Println("Patch NewInsert err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}
