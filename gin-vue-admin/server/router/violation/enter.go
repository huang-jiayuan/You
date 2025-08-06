package violation

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ViolationRouter }

var violationsApi = api.ApiGroupApp.ViolationApiGroup.ViolationApi
