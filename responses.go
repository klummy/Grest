package grest

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON is used to return a JSON response from the API
func ResponseJSON(response http.ResponseWriter, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(data)
}

// NotAllowed Returns an HTTP not allowed method when a user accesses a route that isn't allowed
func NotAllowed(response http.ResponseWriter, errLocation string) {
	response.Header().Set("Content-Type", "application/json")
	http.Error(response, http.StatusText(http.StatusNotImplemented), http.StatusForbidden)
}

// BadRequest for when a user sends a malformed request of any type
func BadRequest(response http.ResponseWriter, err string) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)

	type BadRequestType struct {
		ErrorText      string `json:"error"`
		DefaultMessage string `json:"message"`
	}

	data := BadRequestType{
		ErrorText:      err,
		DefaultMessage: "Bad Request",
	}

	json.NewEncoder(response).Encode(data)
}
