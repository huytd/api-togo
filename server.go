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

// InitDatabase create a new database
func (s *Server) InitDatabase(con string) *gorm.DB {
	DB, _ := gorm.Open("postgres", con)

	if ok := DB.HasTable("maps"); ok {
		fmt.Println("Checking DB.Map OK!")
	} else {
		DB.CreateTable(&models.Map{})
	}

	return DB
}

// InitEngine create a new routing engine
func (s *Server) InitEngine(app *AppController) *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./www", true)))

	api := router.Group("api")
	{
		api.GET("/", app.Index)
		// All routing here
		api.GET("/map", app.MapListing)
		api.GET("/map/:key", app.MapFind)
		api.POST("/map", app.MapCreate)
		api.PUT("/map/:key", app.MapCreate)
		api.DELETE("/map/:key", app.MapDelete)
	}

	return router
}

// Start server, all route come here
func (s *Server) Start(DBHOST string, DBUSER string, DBPWD string, DBNAME string) {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", DBHOST, DBUSER, DBPWD, DBNAME)
	}

	serverPort := ":" + os.Getenv("PORT")
	if serverPort == ":" {
		serverPort = ":8000"
	}

	db := s.InitDatabase(connString)
	app := &AppController{DB: db}
	router := s.InitEngine(app)
	router.Run(serverPort)
}
