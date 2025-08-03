package carouse

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ CarouselImageApi }

var carouselImageService = service.ServiceGroupApp.CarouseServiceGroup.CarouselImageService
