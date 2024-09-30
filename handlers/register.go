package handlers

import (
	"ecommercesite/views"

	"net/http"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) error {

	return Render(w, r, views.Register())

}
