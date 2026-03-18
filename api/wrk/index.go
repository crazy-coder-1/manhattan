package handler

import (
	"encoding/json"
	"manhattan/pkg/utils"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.SuccessResponse("Wrk is Up"))
}