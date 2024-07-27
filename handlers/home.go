package handlers

import (
	"ecommercesite/views/home"

	"ecommercesite/items"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {

	itemList := items.ReturnItems()
	return Render(w, r, home.Index(itemList))

}
