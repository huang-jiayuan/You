package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/admin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gift"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/manage"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/room"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/silence"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/users"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

func PluginInitV2(group *gin.Engine, plugins ...plugin.Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugins[i].Register(group)
	}
}
func bizPluginV2(engine *gin.Engine) {
	PluginInitV2(engine, announcement.Plugin)
	PluginInitV2(engine, room.Plugin)
	PluginInitV2(engine, manage.Plugin)
	PluginInitV2(engine, chamber.Plugin)
	PluginInitV2(engine, silence.Plugin)
	PluginInitV2(engine, users.Plugin)
	PluginInitV2(engine, gift.Plugin)
	PluginInitV2(engine, admin.Plugin)
}
