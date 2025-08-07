package carouse

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ CarouselImageRouter }

var carouselImageApi = api.ApiGroupApp.CarouseApiGroup.CarouselImageApi
