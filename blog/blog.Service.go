package blog

import (
	"blogPost/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewArticle struct {
	Name   string `json:"name" binding:"required"`
	Number int    `json:"number" binding:"required"`
}

type ArticleUpdate struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func GetArticles(c *gin.Context) {
	var articles []model.Article

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&articles).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
	var article model.Article

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func PostArticle(c *gin.Context) {
	var article NewArticle

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newArticle := model.Article{Name: article.Name, Number: article.Number}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newArticle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newArticle)
}

func UpdateArticle(c *gin.Context) {
	var article model.Article

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	var updateArticle ArticleUpdate

	if err := c.ShouldBindJSON(&updateArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&article).Updates(model.Article{Name: updateArticle.Name,
		Number: updateArticle.Number}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func DeleteArticle(c *gin.Context) {
	var article model.Article

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	if err := db.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": "Article deleted"})
}
