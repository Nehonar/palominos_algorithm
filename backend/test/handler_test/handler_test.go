// handlers_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestPingRoute(t *testing.T) {
	// Configurar el servidor de prueba.
	router := mux.NewRouter()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Crear una solicitud para el endpoint de prueba.
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Registrar el handler y enviar la solicitud al servidor de prueba.
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Comprobar el código de estado de la respuesta.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
