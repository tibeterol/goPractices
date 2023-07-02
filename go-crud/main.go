package main

import (
	"github.com/gin-gonic/gin"
	"go-crud/controllers"
	"go-crud/initializers"
)

func init() { // init ozel bir fonk. main icinde ilk once cagriliyor
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/getPosts", controllers.GetPosts)
	r.GET("/getPostsById/:id", controllers.GetPostsById)
	r.PUT("/updatePostById/:id", controllers.UpdatePostById)
	r.DELETE("/deletePostById/:id", controllers.DeletePostById)

	r.Run() // listen and serve on 0.0.0.0:8080
}
