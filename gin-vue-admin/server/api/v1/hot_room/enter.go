package hot_room

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ HostRoomApi }

var hostRoomService = service.ServiceGroupApp.Hot_roomServiceGroup.HostRoomService
