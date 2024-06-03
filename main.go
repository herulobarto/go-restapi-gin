package main

import (
	"github.com/gin-gonic/gin"
	"github.com/herulobarto/go-restapi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnecctDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Index)

	r.Run()
}
