package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"ecommercesite/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	fileServer := http.FileServer(http.Dir("./public"))
	router.Handle("/public/*", http.StripPrefix("/public/", fileServer))

	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/itemPage/{itemID}", (handlers.Make(handlers.HandleItemPage)))
	router.Get("/about", handlers.Make(handlers.HandleAbout))
	router.Get("/register", handlers.Make(handlers.HandleRegister))
	router.Get("/login", handlers.Make(handlers.HandleLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
