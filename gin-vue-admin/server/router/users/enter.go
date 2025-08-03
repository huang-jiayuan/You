package users

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ UserRouter }

var user1Api = api.ApiGroupApp.UsersApiGroup.UserApi
