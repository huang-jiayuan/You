package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/service"

var (
	Api                = new(api)
	serviceRoom        = service.Service.Room
	serviceRoomTagDict = service.Service.RoomTagDict
)

type api struct {
	Room        room
	RoomTagDict roomTagDict
}
