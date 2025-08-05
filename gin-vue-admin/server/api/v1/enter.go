package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/carouse"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/hot_room"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/room_type"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/users"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	CarouseApiGroup   carouse.ApiGroup
	UsersApiGroup     users.ApiGroup
	Hot_roomApiGroup  hot_room.ApiGroup
	Room_typeApiGroup room_type.ApiGroup
}
