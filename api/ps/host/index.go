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

	hostInfo, err := svc.GetHostInfo()
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			500, err.Error(),
		))
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(utils.SuccessResponse(&ps.HostInfoResponse{
		HostName: hostInfo.HostName,
		OS: hostInfo.OS,
		Platform: hostInfo.Platform,
		Uptime: hostInfo.Uptime,
		BootTime: hostInfo.BootTime,
	})))
}