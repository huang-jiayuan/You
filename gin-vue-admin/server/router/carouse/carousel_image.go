package carouse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CarouselImageRouter struct {}

// InitCarouselImageRouter 初始化 carouselImage表 路由信息
func (s *CarouselImageRouter) InitCarouselImageRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	carouselImageRouter := Router.Group("carouselImage").Use(middleware.OperationRecord())
	carouselImageRouterWithoutRecord := Router.Group("carouselImage")
	carouselImageRouterWithoutAuth := PublicRouter.Group("carouselImage")
	{
		carouselImageRouter.POST("createCarouselImage", carouselImageApi.CreateCarouselImage)   // 新建carouselImage表
		carouselImageRouter.DELETE("deleteCarouselImage", carouselImageApi.DeleteCarouselImage) // 删除carouselImage表
		carouselImageRouter.DELETE("deleteCarouselImageByIds", carouselImageApi.DeleteCarouselImageByIds) // 批量删除carouselImage表
		carouselImageRouter.PUT("updateCarouselImage", carouselImageApi.UpdateCarouselImage)    // 更新carouselImage表
	}
	{
		carouselImageRouterWithoutRecord.GET("findCarouselImage", carouselImageApi.FindCarouselImage)        // 根据ID获取carouselImage表
		carouselImageRouterWithoutRecord.GET("getCarouselImageList", carouselImageApi.GetCarouselImageList)  // 获取carouselImage表列表
	}
	{
	    carouselImageRouterWithoutAuth.GET("getCarouselImagePublic", carouselImageApi.GetCarouselImagePublic)  // carouselImage表开放接口
	}
}
