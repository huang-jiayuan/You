package main

import (
	"api/pkg"
	"api/user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	pkg.InitMinio()
	router.Router(r)
	r.Run(":8081")
}
