package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/huytd/api-togo/models"
)

// AppController class
type AppController struct {
	DB *gorm.DB
}

// Index process API root endpoint
func (a *AppController) Index(c *gin.Context) {
	c.JSON(200, gin.H{"status": true})
}

/////////////////////////////////////////////////////

// MapListing is an endpoint to list all map
func (a *AppController) MapListing(c *gin.Context) {
	var all []models.Map
	a.DB.Order("ID desc").Find(&all)
	c.JSON(200, gin.H{"map": all})
}

// MapCreate is an endpoint to add new map
func (a *AppController) MapCreate(c *gin.Context) {
	item := models.Map{Key: c.PostForm("key"), Value: c.PostForm("value")}
	result := a.DB.Create(&item)
	c.JSON(200, gin.H{"result": result})
}

// MapDelete is an endpoint to delete a map
func (a *AppController) MapDelete(c *gin.Context) {
	key := c.Param("key")
	found := models.Map{Key: key}
	result := a.DB.Delete(&found)
	c.JSON(200, gin.H{"result": result})
}

// MapUpdate

// MapFind is an endpoint to get an item with specific key
func (a *AppController) MapFind(c *gin.Context) {
	key := c.Param("key")
	found := models.Map{Key: key}
	c.JSON(200, gin.H{"result": found})
}
