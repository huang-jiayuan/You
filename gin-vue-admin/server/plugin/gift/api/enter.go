package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/service"

var (
	Api                     = new(api)
	serviceGiftInfo         = service.Service.GiftInfo
	serviceUserGiftBackpack = service.Service.UserGiftBackpack
	serviceGiftSendRecord   = service.Service.GiftSendRecord
)

type api struct {
	GiftInfo         giftInfo
	UserGiftBackpack userGiftBackpack
	GiftSendRecord   giftSendRecord
}
