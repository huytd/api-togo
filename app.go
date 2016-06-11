package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/huytd/api-togo/models"
)

// AppController class
type AppController struct {
	db *gorm.DB
}

// Index process API root endpoint
func (a *AppController) Index(c *gin.Context) {
	c.JSON(200, gin.H{"status": true})
}

/////////////////////////////////////////////////////

// MapListing is an endpoint to list all map
func (a *AppController) MapListing(c *gin.Context) {
	var all []models.Map
	a.db.Order("ID desc").Find(&all)
	c.JSON(200, gin.H{"map": all})
}

// MapCreate is an endpoint to add new map
func (a *AppController) MapCreate(c *gin.Context) {
	var item = models.Map{Key: c.PostForm("Key"), Value: c.PostForm("Value")}
	result := a.db.Create(&item)
	c.JSON(200, gin.H{"result": result})
}

// MapDelete is an endpoint to delete a map
func (a *AppController) MapDelete(c *gin.Context) {
	key := c.PostForm("Key")
	var item = models.Map{Key: key}
	result := a.db.Delete(&item)
	c.JSON(200, gin.H{"result": result})
}

// MapUpdate

// MapFind
