package handler

import (
	"encoding/json"
	"manhattan/utils"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string {
		"message": utils.Greet(", Go Wrk is up"),
	})
}