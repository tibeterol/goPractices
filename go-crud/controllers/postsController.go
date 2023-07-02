package controllers

import (
	"github.com/gin-gonic/gin"
	"go-crud/initializers"
	"go-crud/models"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body) // request bodyi kullanmamizi sagliyor

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {

	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, posts)
}

func GetPostsById(c *gin.Context) {

	id := c.Param("id")
	var post models.Post

	initializers.DB.Find(&post, id)

	c.JSON(200, post)
}

func UpdatePostById(c *gin.Context) { // alanlari bodyden id yi queryden aliyoruz
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	var post models.Post

	initializers.DB.Find(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, post)

}

func DeletePostById(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)

}
