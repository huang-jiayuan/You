package violation

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ MuteApi }

var muteService = service.ServiceGroupApp.ViolationServiceGroup.MuteService
