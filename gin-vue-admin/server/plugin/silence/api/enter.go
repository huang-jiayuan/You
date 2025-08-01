package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/silence/service"

var (
	Api         = new(api)
	serviceMute = service.Service.Mute
)

type api struct{ Mute mute }
