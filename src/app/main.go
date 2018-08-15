package main

import (
	"GinWeb/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	r := controllers.RegisterRoutes()
	r.Run(":3000")
}
