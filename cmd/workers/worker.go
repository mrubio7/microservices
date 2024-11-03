package main

import (
	"fmt"
	"ibercs/cmd/workers/matches"
	"ibercs/cmd/workers/players"
	"ibercs/cmd/workers/teams"
	"ibercs/cmd/workers/tournaments"
	"ibercs/pkg/logger"
	"log"
	"net/http"
	"time"
)

func main() {
	logger.Initialize()
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/teams-update", updateTeams)
	http.HandleFunc("/tournaments", findTournaments)
	http.HandleFunc("/find-matches", findMatches)

	logger.Info("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	// Configuramos el header para SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Inicia la rutina para el procesamiento de jugadores
	players.Update(w)
}

func updateTeams(w http.ResponseWriter, r *http.Request) {
	// Configuramos el header para SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Inicia la rutina para el procesamiento de jugadores
	teams.Update(w)
}

func findTournaments(w http.ResponseWriter, r *http.Request) {
	tournaments.Find()
	w.WriteHeader(http.StatusOK)
}

func findHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	// Obtén el tamaño del query parameter "size", con un valor predeterminado de 5000
	size := r.URL.Query().Get("size")
	if size == "" {
		size = "5000" // Valor predeterminado
	}

	// Convierte el tamaño a entero
	var number int
	if _, err := fmt.Sscanf(size, "%d", &number); err != nil {
		http.Error(w, "Invalid size parameter", http.StatusBadRequest)
		return
	}

	// Llama a la función que busca jugadores
	players.Find(number)

	// Calcular el tiempo de ejecución
	elapsedTime := time.Since(startTime)

	// Respuesta con la hora de inicio y el tiempo que tardó
	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf(`{"start_time": "%s", "elapsed_time": "%s"}`, startTime.Format(time.RFC3339), elapsedTime)
	w.Write([]byte(response))
}

func findMatches(w http.ResponseWriter, r *http.Request) {
	matches.Find()
	w.WriteHeader(http.StatusOK)
}
