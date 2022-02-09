package server

import (
	"errors"
	"net/http"
)

var ErrorNotFound = errors.New("server: not found")
var ErrorBadRequest = errors.New("server: bad request")
var ErrorUnauthorized = errors.New("server: unauthorized")

func setErrorMessage(message string) map[string]interface{} {
	m := make(map[string]interface{})

	m["message"] = message

	return m
}

func MakeNotFoundResponse(m string) (int, map[string]interface{}) {
	if m != "" {
		return http.StatusNotFound, setErrorMessage(m)
	}

	return http.StatusNotFound, setErrorMessage("Não encontrado")
}

func MakeBadRequestResponse(m string) (int, map[string]interface{}) {
	if m != "" {
		return http.StatusBadRequest, setErrorMessage(m)
	}

	return http.StatusBadRequest, setErrorMessage("Dados inválidos")
}

func MakeInternalServerErrorResponse() (int, map[string]interface{}) {
	return http.StatusInternalServerError, setErrorMessage("Um erro desconhecido aconteceu")
}
