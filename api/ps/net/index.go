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

	netInfo, err := svc.GetNetworkInfo()
	if err != nil {
		json.NewEncoder(w).Encode(utils.ErrorResponse(
			500, err.Error(),
		))
		return
	}
	netInterfaces := []ps.NetworkInterfaceResponse{}
	for _, intr := range netInfo.Interfaces {
		netInterfaces = append(netInterfaces, ps.NetworkInterfaceResponse{
			Name: intr.Name,
			MTU: intr.MTU,
			HardwareAddress: intr.HardwareAddress,
			Flags: intr.Flags,
			Addresses: intr.Addresses,
		})
	}
	json.NewEncoder(w).Encode(utils.SuccessResponse(&ps.NetworkInfoResponse{
		PublicIP: netInfo.PublicIP,
		Interfaces: netInterfaces,
	}))
}