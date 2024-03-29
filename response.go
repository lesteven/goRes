package goRes

import (
    "net/http"
    "encoding/json"
)

// add content type before status
func header(w http.ResponseWriter, status int) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
}

type Json struct {
    Status int `json:"status"`
    Data interface{} `json:"data"`
}

// takes empty interface b/c JSON data may be represented by different
// structs
func SendJson(w http.ResponseWriter, status int, data interface{}) {
    header(w, status)
    jsonData := Json{status, data}
    json.NewEncoder(w).Encode(jsonData)
}

type ErrJson struct {
    Status int `json:"status"`
    Error string `json:"error"`
}

// error message only takes strings b/c they will all have the same
// json reponse
func SendError(w http.ResponseWriter, status int, errMsg string) {
    header(w, status)
    data := ErrJson {
        Status: status,
        Error: errMsg,
    }
    json.NewEncoder(w).Encode(data)
}
