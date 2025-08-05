package hot_room

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ HotRoomApi }

var hotRoomService = service.ServiceGroupApp.Hot_roomServiceGroup.HotRoomService
