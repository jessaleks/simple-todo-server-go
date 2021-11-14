package utils

import "net/http"

func CheckMethod(r *http.Request, method string) bool {
	return r.Method == method
}
