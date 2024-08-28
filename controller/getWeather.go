package controller

import (
	"defer/utility"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type RepoWeather interface {
	GetCurrentWeather(w http.ResponseWriter, r *http.Request)
}

type repoWeather struct{}

func NewWeather() RepoWeather {
	return &repoWeather{}
}
func (repo *repoWeather) GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	humidity := utility.GenerateRandomNumber(0, 100)
	temperature := utility.GenerateRandomNumber(-20, 250)
	wind := utility.GenerateRandomNumber(0, 40)
	rain := utility.GenerateRandomNumber(0, 40)

	response := map[string]interface{}{
		"humidity":     humidity,
		"temperature":  temperature,
		"wind":         wind,
		"rain":         rain,
		"last_checked": time.Now(),
	}

	log.Printf("type=%v method=%v path=%v", "[INFO]", r.Method, r.URL.Path)
	json.NewEncoder(w).Encode(response)
}
