package app

import (
	"github.com/somkiet073/Golang-api-for-testskill/app/controllers"
)

var server controllers.Server

// Run = run
func Run() {
	// init ค่าให้โปรแกรม connect database mysql
	server.Initialize("mysql")
	server.Run()
}
