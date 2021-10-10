package responseFunctions

import "net/http"

func RespondWithJSON(response http.ResponseWriter, code int, payload []byte) {

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	response.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Authorization, RefreshToken")
	response.WriteHeader(code)
	response.Write(payload)
}
