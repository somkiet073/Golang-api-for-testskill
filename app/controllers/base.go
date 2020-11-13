package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/somkiet073/Golang-api-for-testskill/app/models"
	configs "github.com/somkiet073/Golang-api-for-testskill/app/utils/configs"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
)

// load config file config.yaml
var con configs.Config
var e configs.Env
var conApp, errApp = e.LoadApp()

// Server = server
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize = initialize
func (s *Server) Initialize(dbDriver string) {

	// load config
	c, errcon := con.Load(dbDriver)
	if errcon != nil {
		fmt.Printf("Error load config.")
		log.Fatal("This error:", errcon)
	}

	var err error
	// open connect database
	s.DB, err = gorm.Open(c.Driver, c.Auth)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", c.Driver)
		log.Fatal("This is error:", err)
	} else {
		fmt.Printf("Connect %s database success", c.Driver)
	}

	// this auto migrate
	s.DB.Debug().AutoMigrate(&models.User{}, &models.Project{})

	// set router
	s.Router = mux.NewRouter()

	// init route file routes.go
	s.initializeRoutes()
}

// Run = run
func (s *Server) Run() {

	// load config server port
	var appPort = conApp.ServerPort
	if errApp != nil {
		fmt.Printf("Error load config.")
		log.Fatal("This error:", errApp)
	}

	fmt.Println("Listen to port" + appPort)
	log.Fatal(http.ListenAndServe(appPort, s.Router))
}
