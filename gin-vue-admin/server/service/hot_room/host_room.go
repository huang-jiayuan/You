
package hot_room

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hot_room"
    hot_roomReq "github.com/flipped-aurora/gin-vue-admin/server/model/hot_room/request"
    "gorm.io/gorm"
)

type HostRoomService struct {}
// CreateHostRoom 创建hostRoom表记录
// Author [yourname](https://github.com/yourname)
func (hostRoomService *HostRoomService) CreateHostRoom(ctx context.Context, hostRoom *hot_room.HostRoom) (err error) {
	err = global.GVA_DB.Create(hostRoom).Error
	return err
}

// DeleteHostRoom 删除hostRoom表记录
// Author [yourname](https://github.com/yourname)
func (hostRoomService *HostRoomService)DeleteHostRoom(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hot_room.HostRoom{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&hot_room.HostRoom{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteHostRoomByIds 批量删除hostRoom表记录
// Author [yourname](https://github.com/yourname)
func (hostRoomService *HostRoomService)DeleteHostRoomByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hot_room.HostRoom{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&hot_room.HostRoom{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateHostRoom 更新hostRoom表记录
// Author [yourname](https://github.com/yourname)
func (hostRoomService *HostRoomService)UpdateHostRoom(ctx context.Context, hostRoom hot_room.HostRoom) (err error) {
	err = global.GVA_DB.Model(&hot_room.HostRoom{}).Where("id = ?",hostRoom.ID).Updates(&hostRoom).Error
	return err
}

// GetHostRoom 根据ID获取hostRoom表记录
// Author [yourname](https://github.com/yourname)
func (hostRoomService *HostRoomService)GetHostRoom(ctx context.Context, ID string) (hostRoom hot_room.HostRoom, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&hostRoom).Error
	return
}
// GetHostRoomInfoList 分页获取hostRoom表记录
// Author [yourname](https://github.com/yourname)
func (hostRoomService *HostRoomService)GetHostRoomInfoList(ctx context.Context, info hot_roomReq.HostRoomSearch) (list []hot_room.HostRoom, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hot_room.HostRoom{})
    var hostRooms []hot_room.HostRoom
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
    if info.RoomNum != nil {
        db = db.Where("room_num = ?", *info.RoomNum)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
           orderMap["ID"] = true
           orderMap["CreatedAt"] = true
         	orderMap["room_num"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&hostRooms).Error
	return  hostRooms, total, err
}
func (hostRoomService *HostRoomService)GetHostRoomPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
