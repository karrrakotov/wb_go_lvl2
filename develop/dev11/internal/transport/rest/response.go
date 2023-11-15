package rest

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseOk struct {
	Result interface{} `json:"result"`
}

func ResponseJson(w http.ResponseWriter, status int, structure interface{}) {
	response, err := json.Marshal(structure)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка json.Marshal: " + err.Error(),
		}
		ResponseJson(w, 500, responseError)
		return
	}

	w.WriteHeader(status)
	w.Write(response)
}
