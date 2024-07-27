package handlers

import (
	"ecommercesite/views/itemPage"

	"net/http"
)

func HandleItemPage(w http.ResponseWriter, r *http.Request) error {

	return Render(w, r, itemPage.ItemPage())

}
