
package carouse

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/carouse"
    carouseReq "github.com/flipped-aurora/gin-vue-admin/server/model/carouse/request"
    "gorm.io/gorm"
)

type CarouselImageService struct {}
// CreateCarouselImage 创建carouselImage表记录
// Author [yourname](https://github.com/yourname)
func (carouselImageService *CarouselImageService) CreateCarouselImage(ctx context.Context, carouselImage *carouse.CarouselImage) (err error) {
	err = global.GVA_DB.Create(carouselImage).Error
	return err
}

// DeleteCarouselImage 删除carouselImage表记录
// Author [yourname](https://github.com/yourname)
func (carouselImageService *CarouselImageService)DeleteCarouselImage(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&carouse.CarouselImage{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&carouse.CarouselImage{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCarouselImageByIds 批量删除carouselImage表记录
// Author [yourname](https://github.com/yourname)
func (carouselImageService *CarouselImageService)DeleteCarouselImageByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&carouse.CarouselImage{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&carouse.CarouselImage{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCarouselImage 更新carouselImage表记录
// Author [yourname](https://github.com/yourname)
func (carouselImageService *CarouselImageService)UpdateCarouselImage(ctx context.Context, carouselImage carouse.CarouselImage) (err error) {
	err = global.GVA_DB.Model(&carouse.CarouselImage{}).Where("id = ?",carouselImage.ID).Updates(&carouselImage).Error
	return err
}

// GetCarouselImage 根据ID获取carouselImage表记录
// Author [yourname](https://github.com/yourname)
func (carouselImageService *CarouselImageService)GetCarouselImage(ctx context.Context, ID string) (carouselImage carouse.CarouselImage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&carouselImage).Error
	return
}
// GetCarouselImageInfoList 分页获取carouselImage表记录
// Author [yourname](https://github.com/yourname)
func (carouselImageService *CarouselImageService)GetCarouselImageInfoList(ctx context.Context, info carouseReq.CarouselImageSearch) (list []carouse.CarouselImage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&carouse.CarouselImage{})
    var carouselImages []carouse.CarouselImage
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Image != nil && *info.Image != "" {
        db = db.Where("image = ?", *info.Image)
    }
    if info.Title != nil && *info.Title != "" {
        db = db.Where("title = ?", *info.Title)
    }
    if info.Url != nil && *info.Url != "" {
        db = db.Where("url = ?", *info.Url)
    }
    if info.OrderId != nil {
        db = db.Where("order_id = ?", *info.OrderId)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&carouselImages).Error
	return  carouselImages, total, err
}
func (carouselImageService *CarouselImageService)GetCarouselImagePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
