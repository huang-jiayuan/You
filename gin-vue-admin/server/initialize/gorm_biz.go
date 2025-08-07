package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/carouse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hot_room"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/flipped-aurora/gin-vue-admin/server/model/violation"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(users.User{}, carouse.CarouselImage{}, violation.Violation{}, hot_room.HostRoom{})
	if err != nil {
		return err
	}
	return nil
}
