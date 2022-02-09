package server

import "net/http"

func setErrorMessage(message string) map[string]interface{} {
	m := make(map[string]interface{})

	m["message"] = message

	return m
}

func MakeNotFound(m string) (int, map[string]interface{}) {
	if m != "" {
		return http.StatusNotFound, setErrorMessage(m)
	}

	return http.StatusNotFound, setErrorMessage("Não encontrado")
}

func MakeBadRequest(m string) (int, map[string]interface{}) {
	if m != "" {
		return http.StatusBadRequest, setErrorMessage(m)
	}

	return http.StatusBadRequest, setErrorMessage("Dados inválidos")
}

func MakeInternalServerError() (int, map[string]interface{}) {
	return http.StatusInternalServerError, setErrorMessage("Um erro desconhecido aconteceu")
}
