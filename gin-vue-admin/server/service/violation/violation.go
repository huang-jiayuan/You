
package violation

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/violation"
    violationReq "github.com/flipped-aurora/gin-vue-admin/server/model/violation/request"
    "gorm.io/gorm"
)

type ViolationService struct {}
// CreateViolation 创建violation表记录
// Author [yourname](https://github.com/yourname)
func (violationsService *ViolationService) CreateViolation(ctx context.Context, violations *violation.Violation) (err error) {
	err = global.GVA_DB.Create(violations).Error
	return err
}

// DeleteViolation 删除violation表记录
// Author [yourname](https://github.com/yourname)
func (violationsService *ViolationService)DeleteViolation(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&violation.Violation{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&violation.Violation{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteViolationByIds 批量删除violation表记录
// Author [yourname](https://github.com/yourname)
func (violationsService *ViolationService)DeleteViolationByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&violation.Violation{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&violation.Violation{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateViolation 更新violation表记录
// Author [yourname](https://github.com/yourname)
func (violationsService *ViolationService)UpdateViolation(ctx context.Context, violations violation.Violation) (err error) {
	err = global.GVA_DB.Model(&violation.Violation{}).Where("id = ?",violations.ID).Updates(&violations).Error
	return err
}

// GetViolation 根据ID获取violation表记录
// Author [yourname](https://github.com/yourname)
func (violationsService *ViolationService)GetViolation(ctx context.Context, ID string) (violations violation.Violation, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&violations).Error
	return
}
// GetViolationInfoList 分页获取violation表记录
// Author [yourname](https://github.com/yourname)
func (violationsService *ViolationService)GetViolationInfoList(ctx context.Context, info violationReq.ViolationSearch) (list []violation.Violation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&violation.Violation{})
    var violationss []violation.Violation
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.ViolationMean != nil && *info.ViolationMean != "" {
        db = db.Where("violation_mean = ?", *info.ViolationMean)
    }
    if info.Result != nil && *info.Result != "" {
        db = db.Where("result = ?", *info.Result)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&violationss).Error
	return  violationss, total, err
}
func (violationsService *ViolationService)GetViolationPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
