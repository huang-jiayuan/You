
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/silence/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/silence/model/request"
    "gorm.io/gorm"
)

var Mute = new(mute)

type mute struct {}
// CreateMute 创建mute表记录
// Author [yourname](https://github.com/yourname)
func (s *mute) CreateMute(ctx context.Context, mute *model.Mute) (err error) {
	err = global.GVA_DB.Create(mute).Error
	return err
}

// DeleteMute 删除mute表记录
// Author [yourname](https://github.com/yourname)
func (s *mute) DeleteMute(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.Mute{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.Mute{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMuteByIds 批量删除mute表记录
// Author [yourname](https://github.com/yourname)
func (s *mute) DeleteMuteByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.Mute{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.Mute{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMute 更新mute表记录
// Author [yourname](https://github.com/yourname)
func (s *mute) UpdateMute(ctx context.Context, mute model.Mute) (err error) {
	err = global.GVA_DB.Model(&model.Mute{}).Where("id = ?",mute.ID).Updates(&mute).Error
	return err
}

// GetMute 根据ID获取mute表记录
// Author [yourname](https://github.com/yourname)
func (s *mute) GetMute(ctx context.Context, ID string) (mute model.Mute, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mute).Error
	return
}
// GetMuteInfoList 分页获取mute表记录
// Author [yourname](https://github.com/yourname)
func (s *mute) GetMuteInfoList(ctx context.Context, info request.MuteSearch) (list []model.Mute, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Mute{})
    var mutes []model.Mute
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
	err = db.Find(&mutes).Error
	return  mutes, total, err
}

func (s *mute)GetMutePublic(ctx context.Context) {

}
