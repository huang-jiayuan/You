package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/silence/api"

var (
	Router  = new(router)
	apiMute = api.Api.Mute
)

type router struct{ Mute mute }
