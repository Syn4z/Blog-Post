package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Syn4z/Blog-Post/model"
)

func main() {
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()

	router.GET("/articles", article.GetArticle)
	router.GET("/article/:id", article.GetArticl)
	router.POST("/article", article.PostArticle)
	router.PUT("/article/:id", article.UpdateArticle)
	router.DELETE("/article/:id", article.DeleteArticle)

	log.Fatal(router.Run(":10000"))
}
