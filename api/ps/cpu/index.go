package handler

import (
	"encoding/json"
	"manhattan/pkg/ps"
	"manhattan/pkg/utils"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	svc := ps.NewPSService()

	cpuInfo, err := svc.GetCPUInfo()
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			500, err.Error(),
		))
	}
	res := []*ps.CPUInfoResponse{}
	for _, info := range cpuInfo {
		res = append(res, &ps.CPUInfoResponse{
			Model: info.Model,
			Cores: info.Cores,
			Usage: info.Usage,
		})
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(res))
}