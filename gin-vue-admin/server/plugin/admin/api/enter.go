package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/admin/service"

var (
	Api              = new(api)
	serviceAdminUser = service.Service.AdminUser
)

type api struct{ AdminUser adminUser }
