package routes

import (
	"golang/golang-rest-api/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListRoutes mendaftarkan semua route ke router utama
func ListRoutes(router *gin.Engine) {
	router.GET("/api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test API berhasil",
		})
	})

	router.POST("/api/articles", controllers.StoreArticle)
	router.GET("/api/articles", controllers.IndexArticle)

	router.Run()
}
