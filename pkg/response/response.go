package response

import (
	"encoding/json"
	"net/http"
)

// Retrieve Response
type Retrieve struct {
	Message string `json:"message"`
}

// SendError calls Send by without a count or results
func SendMessage(w http.ResponseWriter, status http.ConnState, message string) {
	var i interface{}
	i = &Retrieve{
		Message: message,
	}
	js, _ := json.Marshal(i)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	w.Write(js)
	return
}

// SendError calls Send by without a count or results
func SendError(w http.ResponseWriter, status http.ConnState, message string) {
	var i interface{}
	i = &Retrieve{
		Message: message,
	}
	js, _ := json.Marshal(i)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	w.Write(js)
	return
}

// SendJSON writes a struct to the writer
func SendJSON(w http.ResponseWriter, status http.ConnState, i interface{}) {
	js, _ := json.Marshal(i)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	if string(js) != "null" {
		w.Write(js)
		return
	}
	return
}
