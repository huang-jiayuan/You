package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/carouse"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/users"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CarouseServiceGroup carouse.ServiceGroup
	UsersServiceGroup   users.ServiceGroup
}
