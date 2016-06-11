package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/huytd/api-togo/models"
)

// Server class
type Server struct {
}

// Start server, all route come here
func (s *Server) Start(DBHOST string, DBUSER string, DBPWD string, DBNAME string) {
	var connString = os.Getenv("DATABASE_URL")
	if connString == "" {
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", DBHOST, DBUSER, DBPWD, DBNAME)
	}
	DB, _ := gorm.Open("postgres", connString)

	if ok := DB.HasTable("posts"); ok {
		fmt.Println("Checking DB.Post OK!")
	} else {
		DB.CreateTable(&models.Post{})
	}

	var router = gin.Default()
	var app = &AppController{db: DB}

	router.Use(static.Serve("/", static.LocalFile("./www", true)))

	api := router.Group("api")
	{
		api.GET("/", app.Index)
		// All routing here
		api.GET("/key", app.PostListing)
		api.POST("/key", app.PostCreate)
		api.DELETE("/key", app.PostDelete)
	}

	// Start HTTP server
	var serverPort = ":" + os.Getenv("PORT")
	if serverPort == ":" {
		serverPort = ":8000"
	}

	router.Run(serverPort)
}
