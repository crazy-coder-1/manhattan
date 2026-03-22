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

	processInfo, err := svc.GetProcessInfo()
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			500, err.Error(),
		))
		return
	}
	res := []*ps.ProcessInfoResponse{}
	for _, info := range processInfo {
		res = append(res, &ps.ProcessInfoResponse{
			PID: info.PID,
			Name: info.Name,
			CPU: info.CPU,
			Mem: info.Mem,
		})
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(res))
}