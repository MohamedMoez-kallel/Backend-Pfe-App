package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Message(message string) map[string]interface{} {

	return map[string]interface{}{"message": message}
}
func Messages(messages string) map[string]interface{} {

	return map[string]interface{}{"error": messages}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Responds(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(data)
}

func Response(w http.ResponseWriter, httpCode int, data interface{}) {
	resp := map[string]interface{}{}
	resp["response"] = data
	response, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}