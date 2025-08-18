
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model/request"
    "gorm.io/gorm"
)

var GiftInfo = new(giftInfo)

type giftInfo struct {}
// CreateGiftInfo 创建giftInfo表记录
// Author [yourname](https://github.com/yourname)
func (s *giftInfo) CreateGiftInfo(ctx context.Context, giftInfo *model.GiftInfo) (err error) {
	err = global.GVA_DB.Create(giftInfo).Error
	return err
}

// DeleteGiftInfo 删除giftInfo表记录
// Author [yourname](https://github.com/yourname)
func (s *giftInfo) DeleteGiftInfo(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.GiftInfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.GiftInfo{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteGiftInfoByIds 批量删除giftInfo表记录
// Author [yourname](https://github.com/yourname)
func (s *giftInfo) DeleteGiftInfoByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.GiftInfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.GiftInfo{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateGiftInfo 更新giftInfo表记录
// Author [yourname](https://github.com/yourname)
func (s *giftInfo) UpdateGiftInfo(ctx context.Context, giftInfo model.GiftInfo) (err error) {
	err = global.GVA_DB.Model(&model.GiftInfo{}).Where("id = ?",giftInfo.ID).Updates(&giftInfo).Error
	return err
}

// GetGiftInfo 根据ID获取giftInfo表记录
// Author [yourname](https://github.com/yourname)
func (s *giftInfo) GetGiftInfo(ctx context.Context, ID string) (giftInfo model.GiftInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&giftInfo).Error
	return
}
// GetGiftInfoInfoList 分页获取giftInfo表记录
// Author [yourname](https://github.com/yourname)
func (s *giftInfo) GetGiftInfoInfoList(ctx context.Context, info request.GiftInfoSearch) (list []model.GiftInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.GiftInfo{})
    var giftInfos []model.GiftInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.GiftName != nil && *info.GiftName != "" {
        db = db.Where("gift_name LIKE ?", "%"+ *info.GiftName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&giftInfos).Error
	return  giftInfos, total, err
}

func (s *giftInfo)GetGiftInfoPublic(ctx context.Context) {

}
