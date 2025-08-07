
package service

import (
	"context"
	"errors"
	"time"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model/request"
    "go.uber.org/zap"
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

// KickFromMic 踢人下麦
// Author [yourname](https://github.com/yourname)
func (s *room) KickFromMic(ctx context.Context, roomId int64, operatorId uint64, targetUserId uint64, reason string) (micPosition int32, err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 1. 验证操作者权限(房主/管理员)
		var member model.RoomMember
		if err := tx.Where("room_id = ? AND user_id = ? AND role IN (1, 2)", roomId, operatorId).First(&member).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("权限不足：只有房主或管理员可以踢人下麦")
			}
			return err
		}

		// 2. 验证目标用户在麦位上
		var roomMic model.RoomMic
		if err := tx.Where("room_id = ? AND user_id = ? AND status = 1", roomId, targetUserId).First(&roomMic).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("目标用户不在麦位上")
			}
			return err
		}

		// 3. 强制释放麦位
		if err := tx.Model(&roomMic).Updates(map[string]interface{}{
			"status":      0,           // 设置为空闲
			"user_id":     nil,         // 清除用户ID
			"occupy_time": nil,         // 清除占用时间
			"updated_by":  operatorId,  // 记录操作者
		}).Error; err != nil {
			return err
		}

		micPosition = int32(roomMic.MicPosition)

		// 4. 记录操作日志
		global.GVA_LOG.Info("踢人下麦操作",
			zap.Int64("roomId", roomId),
			zap.Uint64("operatorId", operatorId),
			zap.Uint64("targetUserId", targetUserId),
			zap.Int32("micPosition", micPosition),
			zap.String("reason", reason),
		)

		return nil
	})

	return micPosition, err
}

// MuteMicUser 禁言管理功能
// Author [yourname](https://github.com/yourname)
func (s *room) MuteMicUser(ctx context.Context, roomId int64, operatorId uint64, targetUserId uint64, action int32, duration int32, reason string) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 1. 验证操作者权限(房主/管理员)
		var operatorMember model.RoomMember
		if err := tx.Where("room_id = ? AND user_id = ? AND role IN (1, 2)", roomId, operatorId).First(&operatorMember).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("权限不足：只有房主或管理员可以进行禁言操作")
			}
			return err
		}

		// 2. 查询目标用户的房间成员记录
		var targetMember model.RoomMember
		err := tx.Where("user_id = ? AND room_id = ?", targetUserId, roomId).First(&targetMember).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 如果目标用户不在房间成员表中，创建一个记录
				targetMember = model.RoomMember{
					RoomId:   roomId,
					UserId:   targetUserId,
					Role:     0, // 普通成员
					IsMuted:  0,
				}
				if err := tx.Create(&targetMember).Error; err != nil {
					return errors.New("创建用户房间成员记录失败")
				}
			} else {
				return errors.New("查询目标用户房间成员记录失败")
			}
		}

		if action == 1 { // 禁言操作
			// 3. 计算禁言结束时间
			var muteEndTime time.Time
			if duration > 0 {
				muteEndTime = time.Now().Add(time.Duration(duration) * time.Minute)
			} else {
				// 如果没有指定时长，默认禁言30分钟
				muteEndTime = time.Now().Add(30 * time.Minute)
			}

			// 4. 更新RoomMember表的禁言状态和结束时间
			if err := tx.Model(&targetMember).Updates(map[string]interface{}{
				"is_muted":      1,
				"mute_end_time": muteEndTime,
				"updated_by":    operatorId,
			}).Error; err != nil {
				return err
			}

			// 5. 如果被禁言用户在麦位上，记录日志
			var roomMic model.RoomMic
			err = tx.Where("room_id = ? AND user_id = ? AND status = 1", roomId, targetUserId).First(&roomMic).Error
			if err == nil {
				global.GVA_LOG.Info("用户在麦位上被禁言",
					zap.Int64("roomId", roomId),
					zap.Uint64("operatorId", operatorId),
					zap.Uint64("targetUserId", targetUserId),
					zap.Int8("micPosition", roomMic.MicPosition),
					zap.Int32("duration", duration),
					zap.String("reason", reason),
				)
			}

			// 6. 记录禁言操作日志
			global.GVA_LOG.Info("禁言操作",
				zap.Int64("roomId", roomId),
				zap.Uint64("operatorId", operatorId),
				zap.Uint64("targetUserId", targetUserId),
				zap.Int32("duration", duration),
				zap.String("muteEndTime", muteEndTime.Format("2006-01-02 15:04:05")),
				zap.String("reason", reason),
			)

		} else if action == 2 { // 解禁操作
			// 7. 检查用户是否被禁言
			if targetMember.IsMuted != 1 {
				return errors.New("用户当前未被禁言")
			}

			// 8. 更新RoomMember表，解除禁言状态
			if err := tx.Model(&targetMember).Updates(map[string]interface{}{
				"is_muted":      0,
				"mute_end_time": nil,
				"updated_by":    operatorId,
			}).Error; err != nil {
				return err
			}

			// 9. 记录解禁操作日志
			global.GVA_LOG.Info("解禁操作",
				zap.Int64("roomId", roomId),
				zap.Uint64("operatorId", operatorId),
				zap.Uint64("targetUserId", targetUserId),
				zap.String("reason", reason),
			)

		} else {
			return errors.New("无效的操作类型，请使用 1-禁言 或 2-解禁")
		}

		return nil
	})

	return err
}
