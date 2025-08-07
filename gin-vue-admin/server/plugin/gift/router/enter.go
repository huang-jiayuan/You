package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/api"

var (
	Router      = new(router)
	apiGiftInfo = api.Api.GiftInfo
)

type router struct{ GiftInfo giftInfo }
