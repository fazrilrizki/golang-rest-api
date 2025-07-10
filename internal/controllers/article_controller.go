package controllers

import (
	"golang/golang-rest-api/internal/models"
	"golang/golang-rest-api/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoreArticle(c *gin.Context) {
	var article models.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&article)
	c.JSON(http.StatusCreated, gin.H{"article": article})
}

func IndexArticle(c *gin.Context) {
	var article []models.Article
	c.JSON(http.StatusOK, gin.H{"articles": article})
}
