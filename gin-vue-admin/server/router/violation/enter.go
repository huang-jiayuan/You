package violation

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ MuteRouter }

var muteApi = api.ApiGroupApp.ViolationApiGroup.MuteApi
