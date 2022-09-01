package helper

import (
	"net/http"
)

func ResponseStatusOkNoData(msg string) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": msg,
	}
}

func ResponseStatusOkWithData(msg string, data interface{}) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": msg,
		"data":    data,
	}
}

func ResponseBadRequest(msg string) (int, map[string]interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"code":    "400",
		"message": msg,
	}
}

func ResponseNotFound(msg string) (int, map[string]interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"code":    "404",
		"message": msg,
	}
}

func ResponseCreateSuccess(msg string) (int, map[string]interface{}) {
	return http.StatusCreated, map[string]interface{}{
		"code":    "201",
		"message": msg,
	}
}

func ResponseForbidden(msg string) (int, map[string]interface{}) {
	return http.StatusForbidden, map[string]interface{}{
		"code":    "403",
		"message": msg,
	}
}
