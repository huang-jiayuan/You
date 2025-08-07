package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/admin/api"

var (
	Router       = new(router)
	apiAdminUser = api.Api.AdminUser
)

type router struct{ AdminUser adminUser }
