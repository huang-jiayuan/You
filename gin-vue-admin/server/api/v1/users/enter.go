package users

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ UserApi }

var user1Service = service.ServiceGroupApp.UsersServiceGroup.UserService
