package controllers

import (
	"encoding/json"
	"io/ioutil"

	"net/http"

	responses "github.com/somkiet073/Golang-api-for-testskill/app/helpers"
	"github.com/somkiet073/Golang-api-for-testskill/app/models"
)

// Login = login
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
	}

	users := models.User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
	}

}
