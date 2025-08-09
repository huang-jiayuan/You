
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model/request"
    "gorm.io/gorm"
)

var UserGiftBackpack = new(userGiftBackpack)

type userGiftBackpack struct {}
// CreateUserGiftBackpack 创建用户背包礼物表记录
// Author [yourname](https://github.com/yourname)
func (s *userGiftBackpack) CreateUserGiftBackpack(ctx context.Context, userGiftBackpack *model.UserGiftBackpack) (err error) {
	err = global.GVA_DB.Create(userGiftBackpack).Error
	return err
}

// DeleteUserGiftBackpack 删除用户背包礼物表记录
// Author [yourname](https://github.com/yourname)
func (s *userGiftBackpack) DeleteUserGiftBackpack(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.UserGiftBackpack{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.UserGiftBackpack{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteUserGiftBackpackByIds 批量删除用户背包礼物表记录
// Author [yourname](https://github.com/yourname)
func (s *userGiftBackpack) DeleteUserGiftBackpackByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.UserGiftBackpack{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.UserGiftBackpack{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateUserGiftBackpack 更新用户背包礼物表记录
// Author [yourname](https://github.com/yourname)
func (s *userGiftBackpack) UpdateUserGiftBackpack(ctx context.Context, userGiftBackpack model.UserGiftBackpack) (err error) {
	err = global.GVA_DB.Model(&model.UserGiftBackpack{}).Where("id = ?",userGiftBackpack.ID).Updates(&userGiftBackpack).Error
	return err
}

// GetUserGiftBackpack 根据ID获取用户背包礼物表记录
// Author [yourname](https://github.com/yourname)
func (s *userGiftBackpack) GetUserGiftBackpack(ctx context.Context, ID string) (userGiftBackpack model.UserGiftBackpack, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&userGiftBackpack).Error
	return
}
// GetUserGiftBackpackInfoList 分页获取用户背包礼物表记录
// Author [yourname](https://github.com/yourname)
func (s *userGiftBackpack) GetUserGiftBackpackInfoList(ctx context.Context, info request.UserGiftBackpackSearch) (list []model.UserGiftBackpack, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.UserGiftBackpack{})
    var userGiftBackpacks []model.UserGiftBackpack
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
	err = db.Find(&userGiftBackpacks).Error
	return  userGiftBackpacks, total, err
}

func (s *userGiftBackpack)GetUserGiftBackpackPublic(ctx context.Context) {

}
