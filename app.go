package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/huytd/api-togo/models"
)

// AppController class
type AppController struct {
	db *gorm.DB
}

func (a *AppController) Index(c *gin.Context) {
	c.JSON(200, gin.H{"status": true})
}

/////////////////////////////////////////////////////

// List all posts
func (a *AppController) PostListing(c *gin.Context) {
	var allPosts []models.Post
	a.db.Order("ID desc").Find(&allPosts)
	c.JSON(200, gin.H{"posts": allPosts})
}

// Add new post
func (a *AppController) PostCreate(c *gin.Context) {
	var post = models.Post{Message: c.PostForm("Message")}
	result := a.db.Create(&post)
	c.JSON(200, gin.H{"result": result})
}

// Delete a post
func (a *AppController) PostDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("ID"))
	var post = models.Post{ID: id}
	result := a.db.Delete(&post)
	c.JSON(200, gin.H{"result": result})
}
