package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) initializeRoutes() {

	group := "/api/v1"
	// Welcome Route
	s.Router.HandleFunc(group+"/", s.welcome).Methods("GET")

	// User Route
	s.Router.HandleFunc(group+"/user", s.createUser).Methods("POST")
	s.Router.HandleFunc(group+"/user/all", s.getAllUser).Methods("GET")
	s.Router.HandleFunc(group+"/user/{id:[1-9]+}", s.updateUser).Methods("PUT")
	s.Router.HandleFunc(group+"/user/{id:[1-9]+}", s.getUserByID).Methods("GET")
	s.Router.HandleFunc(group+"/user/{id:[1-9]+}", s.deleteUser).Methods("DELETE")

	// Project Route
	s.Router.HandleFunc(group+"/project", s.createProject).Methods("POST")
	s.Router.HandleFunc(group+"/project/all", s.getAllProject).Methods("GET")
	s.Router.HandleFunc(group+"/project/{id:[1-9]+}", s.updateProject).Methods("PUT")
	s.Router.HandleFunc(group+"/project/{id:[1-9]+}", s.getProjectByID).Methods("GET")
	s.Router.HandleFunc(group+"/project/{id:[1-9]+}", s.deleteProject).Methods("DELETE")

	myList := s.myRouterName(group)
	s.printRoute(myList)
}


/******************************** Print ***************************************/

func (s *Server) myRouterName(group string) (myList []string) {
	myList = []string{
		" ",
		" =================== api ====================== ",
		" ",
		"-------------------- Home ---------------------",
		"-----------------------------------------------",
		"method[GET] --> " + group + "/",
		" ",
		"-------------------- user ---------------------",
		"-----------------------------------------------",
		"method[POST] --> " + group + "/user",
		"method[GET] --> " + group + "/user/all",
		"method[GET] --> " + group + "/user/{id:[1-9]+}",
		"method[PUT] --> " + group + "/user/{id:[1-9]+}",
		"method[DELETE] --> " + group + "/user/{id:[1-9]+}",
		" ",
		"------------------ project --------------------",
		"-----------------------------------------------",
		"method[POST] --> " + group + "/project",
		"method[GET] --> " + group + "/project/all",
		"method[GET] --> " + group + "/project/{id:[1-9]+}",
		"method[PUT] --> " + group + "/project/{id:[1-9]+}",
		"method[DELETE] --> " + group + "/project/{id:[1-9]+}",
	}
	return
}

func (s *Server) printRoute(list []string) {
	for _, name := range list {
		fmt.Println(name)
	}
}

/************************************** default *************************************************/
func (s *Server) tests(w http.ResponseWriter, r *http.Request) {

	// use json.Marshal
	data := map[string]string{
		"test":  "test",
		"testx": "testx",
	}
	response, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	/** set Header */
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("token", "xxx999xxx")
	/** Write Header */
	w.WriteHeader(http.StatusOK)

	// json marshal
	w.Write(response)

	// WriteString
	// datas := `[{"test":"test"},{"testx":"testx"}]`
	// response with string
	// io.WriteString(w, datas)

	// json NewEncoder
	// datamap := map[string]string{
	// 	"test":  "test",
	// 	"testx": "testx",
	// }
	// json.NewEncoder(w).Encode(datamap)
}
