package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return router
}

func TestHealthRoute(t *testing.T) {
	request, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, response.Code)
	}
}
