package violation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ViolationRouter struct {}

// InitViolationRouter 初始化 violation表 路由信息
func (s *ViolationRouter) InitViolationRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	violationsRouter := Router.Group("violations").Use(middleware.OperationRecord())
	violationsRouterWithoutRecord := Router.Group("violations")
	violationsRouterWithoutAuth := PublicRouter.Group("violations")
	{
		violationsRouter.POST("createViolation", violationsApi.CreateViolation)   // 新建violation表
		violationsRouter.DELETE("deleteViolation", violationsApi.DeleteViolation) // 删除violation表
		violationsRouter.DELETE("deleteViolationByIds", violationsApi.DeleteViolationByIds) // 批量删除violation表
		violationsRouter.PUT("updateViolation", violationsApi.UpdateViolation)    // 更新violation表
	}
	{
		violationsRouterWithoutRecord.GET("findViolation", violationsApi.FindViolation)        // 根据ID获取violation表
		violationsRouterWithoutRecord.GET("getViolationList", violationsApi.GetViolationList)  // 获取violation表列表
	}
	{
	    violationsRouterWithoutAuth.GET("getViolationPublic", violationsApi.GetViolationPublic)  // violation表开放接口
	}
}
