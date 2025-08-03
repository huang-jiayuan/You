
package users

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
    usersReq "github.com/flipped-aurora/gin-vue-admin/server/model/users/request"
    "gorm.io/gorm"
)

type UserService struct {}
// CreateUser 创建user表记录
// Author [yourname](https://github.com/yourname)
func (user1Service *UserService) CreateUser(ctx context.Context, user1 *users.User) (err error) {
	err = global.GVA_DB.Create(user1).Error
	return err
}

// DeleteUser 删除user表记录
// Author [yourname](https://github.com/yourname)
func (user1Service *UserService)DeleteUser(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&users.User{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&users.User{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteUserByIds 批量删除user表记录
// Author [yourname](https://github.com/yourname)
func (user1Service *UserService)DeleteUserByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&users.User{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&users.User{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateUser 更新user表记录
// Author [yourname](https://github.com/yourname)
func (user1Service *UserService)UpdateUser(ctx context.Context, user1 users.User) (err error) {
	err = global.GVA_DB.Model(&users.User{}).Where("id = ?",user1.ID).Updates(&user1).Error
	return err
}

// GetUser 根据ID获取user表记录
// Author [yourname](https://github.com/yourname)
func (user1Service *UserService)GetUser(ctx context.Context, ID string) (user1 users.User, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&user1).Error
	return
}
// GetUserInfoList 分页获取user表记录
// Author [yourname](https://github.com/yourname)
func (user1Service *UserService)GetUserInfoList(ctx context.Context, info usersReq.UserSearch) (list []users.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&users.User{})
    var user1s []users.User
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Username != nil && *info.Username != "" {
        db = db.Where("username = ?", *info.Username)
    }
    if info.Password != nil && *info.Password != "" {
        db = db.Where("password = ?", *info.Password)
    }
    if info.Nickname != nil && *info.Nickname != "" {
        db = db.Where("nickname = ?", *info.Nickname)
    }
    if info.Avatar != nil && *info.Avatar != "" {
        db = db.Where("avatar = ?", *info.Avatar)
    }
    if info.Gender != nil && *info.Gender != "" {
        db = db.Where("gender = ?", *info.Gender)
    }
    if info.VoiceTag != nil && *info.VoiceTag != "" {
        db = db.Where("voice_tag = ?", *info.VoiceTag)
    }
    if info.InterestTags != nil && *info.InterestTags != "" {
        db = db.Where("interest_tags = ?", *info.InterestTags)
    }
    if info.Level != nil {
        db = db.Where("level = ?", *info.Level)
    }
    if info.VipStatus != nil && *info.VipStatus != "" {
        db = db.Where("vip_status = ?", *info.VipStatus)
    }
    if info.VipExpireTime != nil {
        db = db.Where("vip_expire_time = ?", *info.VipExpireTime)
    }
    if info.Balance != nil {
        db = db.Where("balance = ?", *info.Balance)
    }
    if info.Diamond != nil {
        db = db.Where("diamond = ?", *info.Diamond)
    }
    if info.LastLoginIp != nil && *info.LastLoginIp != "" {
        db = db.Where("last_login_ip = ?", *info.LastLoginIp)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
    if info.FreezeReason != nil && *info.FreezeReason != "" {
        db = db.Where("freeze_reason = ?", *info.FreezeReason)
    }
    if info.FreezeEndTime != nil {
        db = db.Where("freeze_end_time = ?", *info.FreezeEndTime)
    }
    if info.IsAuth != nil && *info.IsAuth != "" {
        db = db.Where("is_auth = ?", *info.IsAuth)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&user1s).Error
	return  user1s, total, err
}
func (user1Service *UserService)GetUserPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
