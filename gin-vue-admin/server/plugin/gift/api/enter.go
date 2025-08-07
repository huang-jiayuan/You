package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/service"

var (
	Api             = new(api)
	serviceGiftInfo = service.Service.GiftInfo
)

type api struct{ GiftInfo giftInfo }
