package error

import (
	"net/http"
)

func getStatusDesc(code int) string {
	return map[int]string{
		http.StatusBadRequest:          "Bad Request",
		http.StatusGone:                "Request Aborted/Canceled",
		http.StatusNotFound:            "Not Found",
		http.StatusUnauthorized:        "Unauthorized",
		http.StatusInternalServerError: "Internal Server Error",
		http.StatusOK:                  "OK",
		http.StatusServiceUnavailable:  "Service Unavailable",
	}[code]
}
