package dao

import (
	"context"
	"github.com/CocaineCong/gin-mall/model"
	"gorm.io/gorm"
)

type FavoritesDao struct {
	*gorm.DB
}

func NewFavoritesDao(ctx context.Context) *FavoritesDao {
	return &FavoritesDao{NewDBClient(ctx)}
}

func NewFavoritesDaoByDB(db *gorm.DB) *FavoritesDao {
	return &FavoritesDao{db}
}

// ListFavorites 列出收藏夹的各类信息.
func (dao *FavoritesDao) ListFavorites(id uint) (resp []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=?", id).Find(&resp).Error
	return
}

func (dao *FavoritesDao) FavoriteExistOrNot(pId uint, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id=? AND user_id=?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, nil
}

/*
todo 切片的返回值使用指针传参.
返回值 resp []*model.Favorite 中的 * 表示指向 model.Favorite 结构体的指针。
在 Go 语言中，当我们想要返回一个结构体类型的切片时，通常使用指针切片来提高性能和效率。这是因为切片本身是一个引用类型，它包含了指向底层数组的指针，而不是实际的数据。当我们将切片作为函数的返回值时，避免了将整个切片的数据进行复制的开销，而只需复制切片的指针即可。
在这种情况下，resp []*model.Favorite 表示返回一个指向 model.Favorite 结构体的指针的切片。这意味着返回的切片中的每个元素都是指向一个 model.Favorite 对象的指针。
通过返回指针切片，我们可以在函数调用方修改切片中的数据，而不需要复制整个切片。这在处理大量数据或对性能要求较高的场景中特别有用。
需要注意的是，当返回指针切片时，需要确保在调用该函数后不会修改切片所指向的数据，以避免潜在的并发访问问题。如果需要对返回的切片进行更改，应该在调用方进行适当的拷贝操作。
*/

func (dao *FavoritesDao) CreateFavorite(in *model.Favorite) error {
	return dao.DB.Model(&model.Favorite{}).Create(&in).Error
}

func (dao *FavoritesDao) DeleteFavorite(uId, fId uint) error {
	return dao.DB.Model(&model.Favorite{}).
		Where("id=? AND user_id=?", fId, uId). //todo :由于主键的作用比较强，所以要先判断索引为主键id，同时这是一个软删除.
		Delete(&model.Favorite{}).Error
}
