package hot_room

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ HostRoomRouter }

var hostRoomApi = api.ApiGroupApp.Hot_roomApiGroup.HostRoomApi
