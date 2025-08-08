
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model/request"
    "gorm.io/gorm"
)

var GiftSendRecord = new(giftSendRecord)

type giftSendRecord struct {}
// CreateGiftSendRecord 创建giftSendRecord表记录
// Author [yourname](https://github.com/yourname)
func (s *giftSendRecord) CreateGiftSendRecord(ctx context.Context, giftSendRecord *model.GiftSendRecord) (err error) {
	err = global.GVA_DB.Create(giftSendRecord).Error
	return err
}

// DeleteGiftSendRecord 删除giftSendRecord表记录
// Author [yourname](https://github.com/yourname)
func (s *giftSendRecord) DeleteGiftSendRecord(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.GiftSendRecord{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.GiftSendRecord{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteGiftSendRecordByIds 批量删除giftSendRecord表记录
// Author [yourname](https://github.com/yourname)
func (s *giftSendRecord) DeleteGiftSendRecordByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.GiftSendRecord{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.GiftSendRecord{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateGiftSendRecord 更新giftSendRecord表记录
// Author [yourname](https://github.com/yourname)
func (s *giftSendRecord) UpdateGiftSendRecord(ctx context.Context, giftSendRecord model.GiftSendRecord) (err error) {
	err = global.GVA_DB.Model(&model.GiftSendRecord{}).Where("id = ?",giftSendRecord.ID).Updates(&giftSendRecord).Error
	return err
}

// GetGiftSendRecord 根据ID获取giftSendRecord表记录
// Author [yourname](https://github.com/yourname)
func (s *giftSendRecord) GetGiftSendRecord(ctx context.Context, ID string) (giftSendRecord model.GiftSendRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&giftSendRecord).Error
	return
}
// GetGiftSendRecordInfoList 分页获取giftSendRecord表记录
// Author [yourname](https://github.com/yourname)
func (s *giftSendRecord) GetGiftSendRecordInfoList(ctx context.Context, info request.GiftSendRecordSearch) (list []model.GiftSendRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.GiftSendRecord{})
    var giftSendRecords []model.GiftSendRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&giftSendRecords).Error
	return  giftSendRecords, total, err
}

func (s *giftSendRecord)GetGiftSendRecordPublic(ctx context.Context) {

}
