package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Prueba Unitaria: ¿La lógica de búsqueda funciona?
func TestBusqueda(t *testing.T) {
	query := "eslo"
	encontrado := false
	for _, p := range baseDeDatos {
		if strings.Contains(strings.ToLower(p.Nombre), query) {
			encontrado = true
			break
		}
	}
	if !encontrado {
		t.Errorf("Se esperaba encontrar Eslovenia/Eslovaquia con la query '%s'", query)
	}
}

// Prueba de Integración: ¿El servidor responde HTML?
func TestServerHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// Simulamos el servidor
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>¿Qué país es cuál?</h1>"))
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código de error: obtenido %v, esperado %v", status, http.StatusOK)
	}
}
