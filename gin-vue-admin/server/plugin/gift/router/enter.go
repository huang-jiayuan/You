package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/api"

var (
	Router              = new(router)
	apiUserGiftBackpack = api.Api.UserGiftBackpack
	apiGiftSendRecord   = api.Api.GiftSendRecord
	apiGiftInfo         = api.Api.GiftInfo
)

type router struct {
	UserGiftBackpack userGiftBackpack
	GiftSendRecord   giftSendRecord
	GiftInfo         giftInfo
}
