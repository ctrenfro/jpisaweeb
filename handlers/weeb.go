package handlers

import (
	"jpisaweeb/views"
	"math/rand"
	"net/http"
	"time"
)

func HandleWeeb(w http.ResponseWriter, r *http.Request) error {

	rand.Seed(time.Now().UnixNano())

	//Generate a random integer between 0 and 99
	randomNumber := rand.Intn(5) + 1

	return Render(w, r, views.Weeb(randomNumber))

}
