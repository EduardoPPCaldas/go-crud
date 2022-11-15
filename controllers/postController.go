package controllers

import (
	"net/http"

	"github.com/EduardoPPCaldas/go-crud/initializers"
	"github.com/EduardoPPCaldas/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	var body struct {
		Body string
		Title string
	}

	ctx.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func PostsList(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var body struct {
		Body string
		Title string
	}

	ctx.Bind(&body)
	
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,	
	})

	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	ctx.Status(http.StatusNoContent)
}