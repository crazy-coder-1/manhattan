package handler

import (
	"encoding/json"
	"manhattan/pkg/utils"
	"manhattan/pkg/wrk"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	svc := wrk.NewWrkLoadService()

	query := r.URL.Query()
	endpoint := query.Get("e")
	if endpoint == "" {
		endpoint = "https://example.com"
	}

	method := query.Get("m")
	if method == "" {
		method = "get"
	}

	payloadStr := query.Get("p") // ?p={"key":"value"}

	connectionsStr := query.Get("c")
	if connectionsStr == "" {
		connectionsStr = "10"
	}
	connections, err := strconv.Atoi(connectionsStr)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			http.StatusBadRequest, "invalid connections value",
		))
		return
	}

	durationStr := query.Get("d")
	if durationStr == "" {
		durationStr = "10"
	}
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			http.StatusBadRequest, "Invalid duration value",
		))
		return
	}

	apiDataB := wrk.ApiDataBuilder{}
	apiData := apiDataB.
		SetUrl(endpoint).
		SetMethod(method).
		SetBody(payloadStr).
		Build()

	ctx := r.Context()
	wrkResult, err := svc.Wrk(ctx, apiData, uint(connections), uint(duration))

	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			500, err.Error(),
		))
		return
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(&wrk.WrkResponse{
		Url:             wrkResult.Url,
		Method:          wrkResult.Method,
		TimeTaken:       wrkResult.TimeTaken,
		Connections:     wrkResult.Connections,
		TotalHits:       wrkResult.TotalHits,
		SuccessHits:     wrkResult.SuccessHits,
		FailureHits:     wrkResult.FailureHits,
		SuccessMessages: wrkResult.SuccessMessages,
		FailureMessages: wrkResult.FailureMessages,
	}))
}
