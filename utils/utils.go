package utils

import (
	"io/ioutil"
	"net/http"
)

// ReadRequestBody will return all body data.
func ReadRequestBody(r *http.Request) ([]byte, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
