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

	usageInfo, err := svc.GetUsageInfo()
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			500, err.Error(),
		))
		return
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(&ps.UsageInfoResponse{
		PID: usageInfo.PID,
		RSS: utils.FormatBytes(usageInfo.RSS),
		VMS: utils.FormatBytes(usageInfo.VMS),
		MemoryUsage: usageInfo.MemoryUsage,
		CPUUsage: usageInfo.CPUUsage,
	})) 
}