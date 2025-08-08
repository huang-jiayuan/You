package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/api"

var (
	Router              = new(router)
	apiGiftInfo         = api.Api.GiftInfo
	apiUserGiftBackpack = api.Api.UserGiftBackpack
	apiGiftSendRecord   = api.Api.GiftSendRecord
)

type router struct {
	GiftInfo         giftInfo
	UserGiftBackpack userGiftBackpack
	GiftSendRecord   giftSendRecord
}
