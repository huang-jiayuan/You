package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/gift/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "gift", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(new(model.UserGiftBackpack), new(model.GiftSendRecord), //go:generate go mod tidy
		//go:generate go mod download
		//go:generate go run gen.go

		new(model.GiftInfo),
	)
	g.Execute()
}
