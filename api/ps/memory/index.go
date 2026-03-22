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

	memInfo, err := svc.GetMemoryInfo()
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			500, err.Error(),
		))
		return
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(&ps.MemoryInfoResponse{
		Total: utils.FormatBytes(memInfo.Total), 
		Used: utils.FormatBytes(memInfo.Used), 
		Free: utils.FormatBytes(memInfo.Free), 
		Usage: memInfo.Usage, 
	}))
}