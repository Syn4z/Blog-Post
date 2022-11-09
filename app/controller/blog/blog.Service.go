package blog

import (
	"log"
	"net/http"

	"github.com/Syn4z/Blog-Post/model"

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

func GetArticles(context *gin.Context) {
	var articles []model.Article

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&articles).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, articles)
}

func GetArticle(context *gin.Context) {
	var article model.Article

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", context.Param("id")).First(&article).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	context.JSON(http.StatusOK, article)
}

func PostArticle(context *gin.Context) {
	var article NewArticle

	if err := context.ShouldBindJSON(&article); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newArticle := model.Article{Name: article.Name, Number: article.Number}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newArticle).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, newArticle)
}

func UpdateArticle(context *gin.Context) {
	var article model.Article

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", context.Param("id")).First(&article).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	var updateArticle ArticleUpdate

	if err := context.ShouldBindJSON(&updateArticle); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&article).Updates(model.Article{Name: updateArticle.Name,
		Number: updateArticle.Number}).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, article)
}

func DeleteArticle(context *gin.Context) {
	var article model.Article

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", context.Param("id")).First(&article).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	if err := db.Delete(&article).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"error": "Article deleted"})
}
