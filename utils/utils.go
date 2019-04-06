package utils

import (
	"io/ioutil"
	"net/http"
)

func ReadRequestBody(r *http.Request) []byte {
	b, _ := ioutil.ReadAll(r.Body)

	return b
}
