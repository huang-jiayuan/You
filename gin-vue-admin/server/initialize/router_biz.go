package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		usersRouter := router.RouterGroupApp.Users
		usersRouter.InitUserRouter(privateGroup, publicGroup)
	}
	{
		carouseRouter := router.RouterGroupApp.Carouse
		carouseRouter.InitCarouselImageRouter(privateGroup, publicGroup)
	}
	{
		violationRouter := router.RouterGroupApp.Violation
		violationRouter.InitViolationRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		hot_roomRouter := router.RouterGroupApp.Hot_room
		hot_roomRouter.InitHostRoomRouter(privateGroup, publicGroup)
	}
}
