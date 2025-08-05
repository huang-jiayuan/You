
package hot_room

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hot_room"
    hot_roomReq "github.com/flipped-aurora/gin-vue-admin/server/model/hot_room/request"
    "gorm.io/gorm"
)

type HotRoomService struct {}
// CreateHotRoom 创建hotRoom表记录
// Author [yourname](https://github.com/yourname)
func (hotRoomService *HotRoomService) CreateHotRoom(ctx context.Context, hotRoom *hot_room.HotRoom) (err error) {
	err = global.GVA_DB.Create(hotRoom).Error
	return err
}

// DeleteHotRoom 删除hotRoom表记录
// Author [yourname](https://github.com/yourname)
func (hotRoomService *HotRoomService)DeleteHotRoom(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hot_room.HotRoom{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&hot_room.HotRoom{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteHotRoomByIds 批量删除hotRoom表记录
// Author [yourname](https://github.com/yourname)
func (hotRoomService *HotRoomService)DeleteHotRoomByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hot_room.HotRoom{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&hot_room.HotRoom{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateHotRoom 更新hotRoom表记录
// Author [yourname](https://github.com/yourname)
func (hotRoomService *HotRoomService)UpdateHotRoom(ctx context.Context, hotRoom hot_room.HotRoom) (err error) {
	err = global.GVA_DB.Model(&hot_room.HotRoom{}).Where("id = ?",hotRoom.ID).Updates(&hotRoom).Error
	return err
}

// GetHotRoom 根据ID获取hotRoom表记录
// Author [yourname](https://github.com/yourname)
func (hotRoomService *HotRoomService)GetHotRoom(ctx context.Context, ID string) (hotRoom hot_room.HotRoom, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&hotRoom).Error
	return
}
// GetHotRoomInfoList 分页获取hotRoom表记录
// Author [yourname](https://github.com/yourname)
func (hotRoomService *HotRoomService)GetHotRoomInfoList(ctx context.Context, info hot_roomReq.HotRoomSearch) (list []hot_room.HotRoom, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hot_room.HotRoom{})
    var hotRooms []hot_room.HotRoom
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.RoomId != nil {
        db = db.Where("room_id = ?", *info.RoomId)
    }
    if info.RoomType != nil && *info.RoomType != "" {
        db = db.Where("room_type = ?", *info.RoomType)
    }
    if info.RoomTags != nil && *info.RoomTags != "" {
        db = db.Where("room_tags = ?", *info.RoomTags)
    }
    if info.RoomStatus != nil && *info.RoomStatus != "" {
        db = db.Where("room_status = ?", *info.RoomStatus)
    }
    if info.RoomHost != nil && *info.RoomHost != "" {
        db = db.Where("room_host = ?", *info.RoomHost)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&hotRooms).Error
	return  hotRooms, total, err
}
func (hotRoomService *HotRoomService)GetHotRoomPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
