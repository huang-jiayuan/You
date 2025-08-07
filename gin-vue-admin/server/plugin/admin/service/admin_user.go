
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/admin/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/admin/model/request"
    "gorm.io/gorm"
)

var AdminUser = new(adminUser)

type adminUser struct {}
// CreateAdminUser 创建管理员表记录
// Author [yourname](https://github.com/yourname)
func (s *adminUser) CreateAdminUser(ctx context.Context, adminUser *model.AdminUser) (err error) {
	err = global.GVA_DB.Create(adminUser).Error
	return err
}

// DeleteAdminUser 删除管理员表记录
// Author [yourname](https://github.com/yourname)
func (s *adminUser) DeleteAdminUser(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.AdminUser{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.AdminUser{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAdminUserByIds 批量删除管理员表记录
// Author [yourname](https://github.com/yourname)
func (s *adminUser) DeleteAdminUserByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.AdminUser{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.AdminUser{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAdminUser 更新管理员表记录
// Author [yourname](https://github.com/yourname)
func (s *adminUser) UpdateAdminUser(ctx context.Context, adminUser model.AdminUser) (err error) {
	err = global.GVA_DB.Model(&model.AdminUser{}).Where("id = ?",adminUser.ID).Updates(&adminUser).Error
	return err
}

// GetAdminUser 根据ID获取管理员表记录
// Author [yourname](https://github.com/yourname)
func (s *adminUser) GetAdminUser(ctx context.Context, ID string) (adminUser model.AdminUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&adminUser).Error
	return
}
// GetAdminUserInfoList 分页获取管理员表记录
// Author [yourname](https://github.com/yourname)
func (s *adminUser) GetAdminUserInfoList(ctx context.Context, info request.AdminUserSearch) (list []model.AdminUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AdminUser{})
    var adminUsers []model.AdminUser
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
	err = db.Find(&adminUsers).Error
	return  adminUsers, total, err
}

func (s *adminUser)GetAdminUserPublic(ctx context.Context) {

}
