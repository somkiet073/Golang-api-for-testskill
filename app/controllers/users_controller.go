package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	responses "github.com/somkiet073/Golang-api-for-testskill/app/helpers"
	"github.com/somkiet073/Golang-api-for-testskill/app/models"
)

// ------------------------- CRUD -----------------------------------
func (s *Server) getAllUser(w http.ResponseWriter, r *http.Request) {

	responses.JSON(w, http.StatusOK, "user all")
}

func (s *Server) getUserByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	responses.JSON(w, http.StatusOK, "id: "+id)

}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
	}

	users := models.User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
	}

	responses.JSON(w, http.StatusOK, users)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	responses.JSON(w, http.StatusOK, id)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	responses.JSON(w, http.StatusOK, id)
}
