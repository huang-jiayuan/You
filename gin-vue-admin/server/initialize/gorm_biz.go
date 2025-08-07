package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/carouse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(carouse.CarouselImage{}, users.User{})
	if err != nil {
		return err
	}
	return nil
}
