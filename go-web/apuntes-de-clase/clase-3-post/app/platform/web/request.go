package web

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidContentType = errors.New("invalid content type")
)

func RequestJSON(r *http.Request, ptr any) (err error) {
	if r.Header.Get("Content-Type") != "application/json" {
		err = ErrInvalidContentType
		return
	}

	err = json.NewDecoder(r.Body).Decode(ptr)
	if err != nil {
		println(err.Error())
		return err
	}

	return
}
