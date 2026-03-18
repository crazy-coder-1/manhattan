package handler

import (
	"encoding/json"
	"manhattan/pkg/ps"
	"manhattan/pkg/utils"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

    svc := ps.NewPSService()
    handler := ps.NewPSHandler(svc)

	parts := strings.Split(r.URL.Path, "/")
	// /api/ps/cpu → ["", "api", "ps", "cpu"]

    if len(parts) == 3 {
        json.NewEncoder(w).Encode(utils.SuccessResponse("PS is Up"))
        return
    }

	if len(parts) != 4 {
        json.NewEncoder(w).Encode(utils.ErrorResponse(404, "Invalid Url, not found"))
        return
	}

    switch parts[3] {
    case "cpu":
        json.NewEncoder(w).Encode(handler.CPUInfo())
    case "memory":
        json.NewEncoder(w).Encode(handler.MemoryInfo())
    case "host":
        json.NewEncoder(w).Encode(handler.HostInfo())
    case "process":
        json.NewEncoder(w).Encode(handler.ProcessInfo())
    case "usage":
        json.NewEncoder(w).Encode(handler.UsageInfo())
    default:
        json.NewEncoder(w).Encode(utils.SuccessResponse("Invalid Url, not found"))
        return
    }

	// json.NewEncoder(w).Encode(utils.SuccessResponse("PS is Up"))

    // switch r.Method {

    // case http.MethodGet:
    //     // GET /api/users
    //     json.NewEncoder(w).Encode(map[string]string{
    //         "action": "get users",
    //     })

    // case http.MethodPost:
    //     // POST /api/users
    //     json.NewEncoder(w).Encode(map[string]string{
    //         "action": "create user",
    //     })

    // case http.MethodPut:
    //     // PUT /api/users?id=1
    //     json.NewEncoder(w).Encode(map[string]string{
    //         "action": "update user",
    //     })

    // case http.MethodDelete:
    //     // DELETE /api/users?id=1
    //     json.NewEncoder(w).Encode(map[string]string{
    //         "action": "delete user",
    //     })

    // default:
    //     http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    // }
	// json.NewEncoder(w).Encode(utils.SuccessResponse("PS is Up"))
}