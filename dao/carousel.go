package dao

import (
	"context"
	"github.com/CocaineCong/gin-mall/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

//该函数的目的是根据给定的 gorm.DB 对象创建一个新的 UserDao 实例，并将该对象的指针作为结果返回。

// GetCarouselById 根据id获取Carousel
func (dao *CarouselDao) GetCarouselById(id uint) (carousel *model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Where("id=?", id).First(&carousel).Error
	return
}

// ListCarousel 用于展示
func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
