package repository

import (
	"boardType/internal/domain"
	"boardType/internal/page"
	"boardType/internal/repository/model"
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

func (r BoardTypeRepository) Create(ctx context.Context, btd domain.BoardType) error {
	m := model.ToModel(btd)
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
		return []domain.BoardType{}, 0, errors.New("internal server error")
	}
	return model.ToDomainLIst(models), cnt, nil
}

func (r BoardTypeRepository) Delete(ctx context.Context, cafeId int, typeId int) error {
	var m model.BoardType
	_, err := r.db.NewDelete().Model(&m).Where("id =? and cafe_id = ?", typeId, cafeId).Exec(ctx)
	if err != nil {
		log.Println("Delete NewDelete err: ", err)
		return errors.New("internal server error")
	}
	return nil
}

func (r BoardTypeRepository) Patch(
	ctx context.Context, cafeId int, id int,
	validFunc func(domains []domain.BoardType) (domain.BoardType, error),
	mergeFunc func(filtered domain.BoardType) domain.BoardType) error {
	var models []model.BoardType
	err := r.db.NewSelect().Model(&models).Where("id =? and cafe_id = ?", id, cafeId).Scan(ctx)
	if err != nil {
		log.Println("Patch NewSelect err: ", err)
		return errors.New("internal server error")
	}
	validD, err := validFunc(model.ToDomainLIst(models))
	if err != nil {
		return err
	}
	mergedDomain := mergeFunc(validD)
	mergedModel := model.ToModel(mergedDomain)
	_, err = r.db.NewInsert().Model(&mergedModel).On("CONFLICT (id) DO UPDATE").Exec(ctx)
	if err != nil {
		log.Println("Patch NewInsert err: ", err)
		return errors.New("internal server error")
	}
	return nil
}
