package handlers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitRoutes() {
	router := mux.NewRouter()

	// Static files from 'public' folder
	staticFileDirectory := http.Dir("./public")
	staticFileHandler := http.StripPrefix("/index", http.FileServer(staticFileDirectory))
	router.PathPrefix("/index").Handler(staticFileHandler).Methods("GET")

	// *** API routers ***
	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// CORS config
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	http.ListenAndServe(":"+PORT, handler)
}
