package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/service"

var (
	Api                     = new(api)
	serviceUserGiftBackpack = service.Service.UserGiftBackpack
	serviceGiftSendRecord   = service.Service.GiftSendRecord
	serviceGiftInfo         = service.Service.GiftInfo
)

type api struct {
	UserGiftBackpack userGiftBackpack
	GiftSendRecord   giftSendRecord
	GiftInfo         giftInfo
}
