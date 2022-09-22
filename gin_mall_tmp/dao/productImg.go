package dao

import (
	"context"
	"gin_mall_tmp/model"

	"gorm.io/gorm"
)

type productImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *productImgDao {
	return &productImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *productImgDao {
	return &productImgDao{db}
}

func (dao *productImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}
