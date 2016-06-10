package main

import (
	"fmt"
	"net/http"
	"os"

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
	var connString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", DBHOST, DBUSER, DBPWD, DBNAME)
	DB, _ := gorm.Open("postgres", connString)

	if ok := DB.HasTable("posts"); ok {
		fmt.Println("Checking DB.Post OK!")
	} else {
		DB.CreateTable(&models.Post{})
	}

	var router = gin.Default()
	var app = &AppController{db: DB}

	router.Static("/css", "./public/css")
	router.LoadHTMLGlob("public/*.html")

	// All routing here
	router.GET("/", app.Home)
	router.POST("/add", app.Add)
	router.POST("/delete", app.Delete)

	// Start HTTP server
	var serverPort = ":" + os.Getenv("PORT")
	if serverPort == ":" {
		serverPort = ":8000"
	}
	http.ListenAndServe(serverPort, router)
}
