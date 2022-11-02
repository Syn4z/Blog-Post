package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Syn4z/Blog-Post/blog"
	"github.com/Syn4z/Blog-Post/model"
)

func main() {
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()

	router.GET("/articles", blog.GetArticles)
	router.GET("/article/:id", blog.GetArticle)
	router.POST("/article", blog.PostArticle)
	router.PUT("/article/:id", blog.UpdateArticle)
	router.DELETE("/article/:id", blog.DeleteArticle)

	log.Fatal(router.Run(":10000"))
}
