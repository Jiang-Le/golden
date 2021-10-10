package routers

import "github.com/gin-gonic/gin"

func init() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	v1.GET("/albums", )
}
