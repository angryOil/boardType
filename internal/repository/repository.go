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
	cnt, err := r.db.NewSelect().Model(&models).ColumnExpr("name,id,substring(bt.description,1,50) as description").Where("cafe_id = ?", cafeId).Limit(reqPage.Size).Offset(reqPage.OffSet).Order("id desc").ScanAndCount(ctx)
	if err != nil {
		log.Println("GetListByCafeId NewSelect err: ", err)
		return []domain.BoardType{}, 0, errors.New("internal server error")
	}
	return model.ToDomainLIst(models), cnt, nil
}
