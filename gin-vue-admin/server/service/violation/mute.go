
package violation

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/violation"
    violationReq "github.com/flipped-aurora/gin-vue-admin/server/model/violation/request"
    "gorm.io/gorm"
)

type MuteService struct {}
// CreateMute 创建mute表记录
// Author [yourname](https://github.com/yourname)
func (muteService *MuteService) CreateMute(ctx context.Context, mute *violation.Mute) (err error) {
	err = global.GVA_DB.Create(mute).Error
	return err
}

// DeleteMute 删除mute表记录
// Author [yourname](https://github.com/yourname)
func (muteService *MuteService)DeleteMute(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&violation.Mute{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&violation.Mute{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMuteByIds 批量删除mute表记录
// Author [yourname](https://github.com/yourname)
func (muteService *MuteService)DeleteMuteByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&violation.Mute{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&violation.Mute{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMute 更新mute表记录
// Author [yourname](https://github.com/yourname)
func (muteService *MuteService)UpdateMute(ctx context.Context, mute violation.Mute) (err error) {
	err = global.GVA_DB.Model(&violation.Mute{}).Where("id = ?",mute.ID).Updates(&mute).Error
	return err
}

// GetMute 根据ID获取mute表记录
// Author [yourname](https://github.com/yourname)
func (muteService *MuteService)GetMute(ctx context.Context, ID string) (mute violation.Mute, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mute).Error
	return
}
// GetMuteInfoList 分页获取mute表记录
// Author [yourname](https://github.com/yourname)
func (muteService *MuteService)GetMuteInfoList(ctx context.Context, info violationReq.MuteSearch) (list []violation.Mute, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&violation.Mute{})
    var mutes []violation.Mute
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.UserId != nil && *info.UserId != "" {
        db = db.Where("user_id = ?", *info.UserId)
    }
    if info.MuteType != nil && *info.MuteType != "" {
        db = db.Where("mute_type = ?", *info.MuteType)
    }
    if info.Reason != nil && *info.Reason != "" {
        db = db.Where("reason = ?", *info.Reason)
    }
    if info.OperatorId != nil {
        db = db.Where("operator_id = ?", *info.OperatorId)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
    if info.MuteDay != nil && *info.MuteDay != "" {
        db = db.Where("mute_day = ?", *info.MuteDay)
    }
    if info.MuteResult != nil && *info.MuteResult != "" {
        db = db.Where("mute_result = ?", *info.MuteResult)
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
func (muteService *MuteService)GetMutePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
