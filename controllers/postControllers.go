package controllers

import (
	initializers "goCRUD/Initializers"
	"goCRUD/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//get data from req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get the id from url
	id := c.Param("id")
	//get the posts
	var posts models.Post
	initializers.DB.First(&posts, id)
	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id from url
	id := c.Param("id")
	// get data from req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//find the post we're updating
	var posts models.Post
	initializers.DB.First(&posts, id)
	//Update it
	initializers.DB.Model(&posts).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	//repond with it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsDelete(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")
	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	//repond with it
	c.Status(200)
}
