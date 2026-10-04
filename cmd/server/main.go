package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"service":"md-intake-lakehouse","status":"UP"}`))
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Servicio md-intake-lakehouse iniciado en puerto 8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
