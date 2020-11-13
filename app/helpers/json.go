package helpers

import (
	"encoding/json"
	"net/http"
	
)

func JSON(w http.ResponseWriter, code int, payload interface{}){
	response, err := json.Marshal(payload)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}