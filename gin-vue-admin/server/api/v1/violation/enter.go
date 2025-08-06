package violation

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ViolationApi }

var violationsService = service.ServiceGroupApp.ViolationServiceGroup.ViolationService
