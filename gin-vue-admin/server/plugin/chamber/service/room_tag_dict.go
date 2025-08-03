
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model/request"
    "gorm.io/gorm"
)

var RoomTagDict = new(roomTagDict)

type roomTagDict struct {}
// CreateRoomTagDict 创建房间话题表记录
// Author [yourname](https://github.com/yourname)
func (s *roomTagDict) CreateRoomTagDict(ctx context.Context, roomTagDict *model.RoomTagDict) (err error) {
	err = global.GVA_DB.Create(roomTagDict).Error
	return err
}

// DeleteRoomTagDict 删除房间话题表记录
// Author [yourname](https://github.com/yourname)
func (s *roomTagDict) DeleteRoomTagDict(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.RoomTagDict{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.RoomTagDict{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteRoomTagDictByIds 批量删除房间话题表记录
// Author [yourname](https://github.com/yourname)
func (s *roomTagDict) DeleteRoomTagDictByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.RoomTagDict{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.RoomTagDict{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateRoomTagDict 更新房间话题表记录
// Author [yourname](https://github.com/yourname)
func (s *roomTagDict) UpdateRoomTagDict(ctx context.Context, roomTagDict model.RoomTagDict) (err error) {
	err = global.GVA_DB.Model(&model.RoomTagDict{}).Where("id = ?",roomTagDict.ID).Updates(&roomTagDict).Error
	return err
}

// GetRoomTagDict 根据ID获取房间话题表记录
// Author [yourname](https://github.com/yourname)
func (s *roomTagDict) GetRoomTagDict(ctx context.Context, ID string) (roomTagDict model.RoomTagDict, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&roomTagDict).Error
	return
}
// GetRoomTagDictInfoList 分页获取房间话题表记录
// Author [yourname](https://github.com/yourname)
func (s *roomTagDict) GetRoomTagDictInfoList(ctx context.Context, info request.RoomTagDictSearch) (list []model.RoomTagDict, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.RoomTagDict{})
    var roomTagDicts []model.RoomTagDict
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
	err = db.Find(&roomTagDicts).Error
	return  roomTagDicts, total, err
}

func (s *roomTagDict)GetRoomTagDictPublic(ctx context.Context) {

}
