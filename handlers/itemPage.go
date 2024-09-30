package handlers

import (
	"ecommercesite/items"
	"ecommercesite/views"
	"net/http"
)

func HandleItemPage(w http.ResponseWriter, r *http.Request) error {
	itemIdString := r.PathValue("itemID")

	item := items.ReturnItem(itemIdString)

	return Render(w, r, views.ItemPage(item))

}
