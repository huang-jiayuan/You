package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/carouse"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/users"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Carouse carouse.RouterGroup
	Users   users.RouterGroup
}
