package controllers

import (
	"encoding/json"
	"errors"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/somkiet073/Golang-api-for-testskill/app/auth"

	"github.com/gorilla/mux"

	responses "github.com/somkiet073/Golang-api-for-testskill/app/helpers"
	"github.com/somkiet073/Golang-api-for-testskill/app/models"
)

// ------------------------- CRUD -----------------------------------

func (s *Server) getAllProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project

	if _, err := auth.ExtractTokenID(r); err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	projects, err := project.FindAllProject(s.DB)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w)(http.StatusOK, projects)
}

func (s *Server) getProjectByID(w http.ResponseWriter, r *http.Request) {

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

	var project models.Project
	projectRs, err := project.FindProjectByID(s.DB, id)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w)(http.StatusOK, projectRs)
}

func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	body, err := responses.Bind(r)
	if err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
		return
	}

	var projects models.Project
	if err = json.Unmarshal(body, &projects); err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
		return
	}

	// Prepare
	projects.ID = 0
	projects.Title = html.EscapeString(strings.TrimSpace(projects.Title))
	projects.Description = html.EscapeString(strings.TrimSpace(projects.Description))
	projects.User = models.User{}
	projects.CreatedAt = time.Now()
	projects.UpdatedAt = time.Now()

	//err = Validate

	id, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if id != projects.UserID {
		responses.JSON(w)(http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	projectCreate, err := projects.CreateProject(s.DB)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err.Error())
		return
	}

	responses.JSON(w)(http.StatusCreated, projectCreate)
}

func (s *Server) updateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	var project models.Project
	err = s.DB.Debug().Model(&models.Project{}).Where("id=?", id).Take(&project).Error
	if err != nil {
		responses.JSON(w)(http.StatusNoContent, errors.New("Post not found"))
		return
	}

	body, err := responses.Bind(r)
	if err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
		return
	}

	var projectsUpdate models.Project
	if err = json.Unmarshal(body, &projectsUpdate); err != nil {
		responses.JSON(w)(http.StatusUnprocessableEntity, err)
		return
	}

	// Prepare
	projectsUpdate.ID = project.ID
	projectsUpdate.Title = html.EscapeString(strings.TrimSpace(projectsUpdate.Title))
	projectsUpdate.Description = html.EscapeString(strings.TrimSpace(projectsUpdate.Description))
	projectsUpdate.User = models.User{}
	projectsUpdate.UpdatedAt = time.Now()

	if _, err := auth.ExtractTokenID(r); err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	updateProjects, err := projectsUpdate.UpdateProject(s.DB)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w)(http.StatusOK, updateProjects)
}

func (s *Server) deleteProject(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.JSON(w)(http.StatusInternalServerError, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.JSON(w)(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	projects := models.Project{}
	err = s.DB.Debug().Model(&models.Project{}).Where("id=?", id).Take(&projects).Error
	if err != nil {
		responses.JSON(w)(http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	if uid != projects.UserID {
		responses.JSON(w)(http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = projects.DeleteProject(s.DB, id, uid)
	if err != nil {
		responses.JSON(w)(http.StatusBadRequest, err)
		return
	}
	responses.JSON(w)(http.StatusNoContent, "")
}
