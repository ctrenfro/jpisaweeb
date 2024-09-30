package handlers

import (
	"ecommercesite/views"

	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {

	return Render(w, r, views.Login())

}
