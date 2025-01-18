package handlers

import (
	"jpisaweeb/views"

	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {

	return Render(w, r, views.Index())

}
