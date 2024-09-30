package handlers

import (
	"ecommercesite/views"

	"net/http"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) error {

	return Render(w, r, views.About())

}
