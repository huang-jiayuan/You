package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chamber/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "chamber", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(new(model.Room), //go:generate go mod tidy
		//go:generate go mod download
		//go:generate go run gen.go

		new(model.RoomTagDict),
	)
	g.Execute()
}
