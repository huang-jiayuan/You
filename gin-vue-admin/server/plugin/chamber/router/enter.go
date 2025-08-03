package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/api"

var (
	Router         = new(router)
	apiRoom        = api.Api.Room
	apiRoomTagDict = api.Api.RoomTagDict
)

type router struct {
	Room        room
	RoomTagDict roomTagDict
}
