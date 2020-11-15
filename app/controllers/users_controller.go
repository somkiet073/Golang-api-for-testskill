package controllers

import (
	"encoding/json"
	"errors"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/somkiet073/Golang-api-for-testskill/app/auth"
	helper "github.com/somkiet073/Golang-api-for-testskill/app/helpers"
	responses "github.com/somkiet073/Golang-api-for-testskill/app/helpers"
	"github.com/somkiet073/Golang-api-for-testskill/app/models"
)

// ------------------------- CRUD -----------------------------------
func (s *Server) getAllUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if _, err := auth.ExtractTokenID(r); err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	users, err := user.FindAllUser(s.DB)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w)(http.StatusOK, users)
}

func (s *Server) getUserByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	if _, err := auth.ExtractTokenID(r); err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	var user models.User
	users, err := user.FindUserByID(s.DB, id)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w)(http.StatusOK, users)

}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {

	// bind payload
	body, err := responses.Bind(r)
	if err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
		return
	}

	var users models.User
	err = json.Unmarshal(body, &users)
	if err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
		return
	}

	// prepare
	hashpass, err := helper.Hash(users.Password)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}
	users.ID = 0
	users.Fristname = html.EscapeString(strings.TrimSpace(users.Fristname))
	users.Lastname = html.EscapeString(strings.TrimSpace(users.Lastname))
	users.Nickname = html.EscapeString(strings.TrimSpace(users.Nickname))
	users.Email = html.EscapeString(strings.TrimSpace(users.Email))
	users.Password = string(hashpass)
	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	if _, err := auth.ExtractTokenID(r); err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	usersCreate, err := users.CreateUser(s.DB)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err.Error())
		return
	}

	responses.JSON(w)(http.StatusCreated, usersCreate)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	// bind payload
	body, err := responses.Bind(r)
	if err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
		return
	}

	var users models.User
	err = json.Unmarshal(body, &users)
	if err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
		return
	}
	// prepare
	users.Fristname = html.EscapeString(strings.TrimSpace(users.Fristname))
	users.Lastname = html.EscapeString(strings.TrimSpace(users.Lastname))
	users.Nickname = html.EscapeString(strings.TrimSpace(users.Nickname))
	users.Email = html.EscapeString(strings.TrimSpace(users.Email))
	users.UpdatedAt = time.Now()

	exid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	exids := uint64(exid)
	if exids != id {
		responses.JSON(w)(http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	usersUpdate, err := users.UpdateUser(s.DB, id)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w)(http.StatusOK, usersUpdate)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	if _, err := auth.ExtractTokenID(r); err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	var users models.User
	_, err = users.DeleteUser(s.DB, id)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w)(http.StatusOK, "")
}
