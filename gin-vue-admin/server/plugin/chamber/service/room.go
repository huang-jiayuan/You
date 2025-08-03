
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model/request"
    "gorm.io/gorm"
)

var Room = new(room)

type room struct {}
// CreateRoom 创建房间表记录
// Author [yourname](https://github.com/yourname)
func (s *room) CreateRoom(ctx context.Context, room *model.Room) (err error) {
	err = global.GVA_DB.Create(room).Error
	return err
}

// DeleteRoom 删除房间表记录
// Author [yourname](https://github.com/yourname)
func (s *room) DeleteRoom(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.Room{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.Room{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteRoomByIds 批量删除房间表记录
// Author [yourname](https://github.com/yourname)
func (s *room) DeleteRoomByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.Room{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.Room{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateRoom 更新房间表记录
// Author [yourname](https://github.com/yourname)
func (s *room) UpdateRoom(ctx context.Context, room model.Room) (err error) {
	err = global.GVA_DB.Model(&model.Room{}).Where("id = ?",room.ID).Updates(&room).Error
	return err
}

// GetRoom 根据ID获取房间表记录
// Author [yourname](https://github.com/yourname)
func (s *room) GetRoom(ctx context.Context, ID string) (room model.Room, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&room).Error
	return
}
// GetRoomInfoList 分页获取房间表记录
// Author [yourname](https://github.com/yourname)
func (s *room) GetRoomInfoList(ctx context.Context, info request.RoomSearch) (list []model.Room, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Room{})
    var rooms []model.Room
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
	err = db.Find(&rooms).Error
	return  rooms, total, err
}

func (s *room)GetRoomPublic(ctx context.Context) {

}
