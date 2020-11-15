package controllers

import (
	"encoding/json"
	"html"
	"strings"
	"time"

	"github.com/somkiet073/Golang-api-for-testskill/app/auth"

	"net/http"

	helper "github.com/somkiet073/Golang-api-for-testskill/app/helpers"
	responses "github.com/somkiet073/Golang-api-for-testskill/app/helpers"
	"github.com/somkiet073/Golang-api-for-testskill/app/models"
)

// Login = login
func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	body, err := responses.Bind(r)
	if err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
	}

	var users models.User
	if err = json.Unmarshal(body, &users); err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
	}

	// prepare
	users.ID = 0
	users.Fristname = html.EscapeString(strings.TrimSpace(users.Fristname))
	users.Lastname = html.EscapeString(strings.TrimSpace(users.Lastname))
	users.Nickname = html.EscapeString(strings.TrimSpace(users.Nickname))
	users.Email = html.EscapeString(strings.TrimSpace(users.Email))
	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	// validate
	token, err := s.validateLogin(users.Email, users.Password)
	if err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err.Error())
		return
	}
	responses.JSON(w)(http.StatusOK, token)
}

// ValidateLogin = validateLogin
func (s *Server) validateLogin(email string, password string) (string, error) {

	var err error
	var users models.User

	err = s.DB.Debug().Model(&models.User{}).Where("email=? ", email).Take(&users).Error
	if err != nil {
		return "", err
	}

	// // Compare password
	err = helper.VerifyPassword([]byte(users.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// create token
	return auth.CreateToken(users.ID)
}
