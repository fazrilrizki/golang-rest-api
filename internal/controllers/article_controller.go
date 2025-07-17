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

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}

	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	article.Title = input.Title
	article.Content = input.Content

	database.DB.Save(&article)

	c.JSON(http.StatusOK, gin.H{"article": article})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}

	database.DB.Delete(&article)

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
