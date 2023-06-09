package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegesterController() {

	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {

	enc := json.NewDecoder(w)
	enc.Encode(data)

}
