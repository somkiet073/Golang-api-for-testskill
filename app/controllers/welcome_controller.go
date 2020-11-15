package controllers

import (
	"net/http"

	"github.com/somkiet073/Golang-api-for-testskill/app/helpers"
)

func (server *Server) welcome(w http.ResponseWriter, r *http.Request) {
	helpers.JSON(w)(http.StatusOK, "Welcome to api")
}
