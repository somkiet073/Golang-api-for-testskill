package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Bind = bind
func Bind(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

// JSON = json
func JSON(w http.ResponseWriter) func(int, interface{}) {
	return func(code int, payload interface{}) {
		response, err := json.Marshal(payload)
		if err != nil {
			panic(err.Error())
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(code)
		w.Write(response)
	}
}

// // JSON = jSON
// func JSON(w http.ResponseWriter, code int, payload interface{}) {
// 	response, err := json.Marshal(payload)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	w.Header().Set("content-type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }
