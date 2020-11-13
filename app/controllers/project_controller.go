package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/somkiet073/Golang-api-for-testskill/app/auth"

	"github.com/gorilla/mux"

	responses "github.com/somkiet073/Golang-api-for-testskill/app/helpers"
	"github.com/somkiet073/Golang-api-for-testskill/app/models"
)

// ------------------------- CRUD -----------------------------------

func (s *Server) getAllProject(w http.ResponseWriter, r *http.Request) {
	project := models.Project{}

	projects, err := project.FindAllProject(s.DB)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, projects)
}

func (s *Server) getProjectByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
	}
	project := models.Project{}

	projectRs, err := project.FindProjectByID(s.DB, id)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, projectRs)

}

func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
	}

	projects := models.Project{}
	err = json.Unmarshal(body, &projects)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	projects.Prepare()
	//err = projects.Validate()

	id, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.JSON(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if id != projects.UserID {
		responses.JSON(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	projectCreate, err := projects.CreateProject(s.DB)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	responses.JSON(w, http.StatusCreated, projectCreate)
}

func (s *Server) updateProject(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	responses.JSON(w, http.StatusOK, id)
}

func (s *Server) deleteProject(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.JSON(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	projects := models.Project{}
	err = s.DB.Debug().Model(&models.Project{}).Where("id=?", id).Take(&projects).Error
	if err != nil {
		responses.JSON(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	if uid != projects.UserID {
		responses.JSON(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = projects.DeleteProject(s.DB, id, uid)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, "")
}
