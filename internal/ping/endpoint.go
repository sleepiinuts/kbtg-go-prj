package ping

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.Write([]byte("pong"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %v", err)
		}
	}()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// expect json body
	var req struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Address struct {
			No   int    `json:"no"`
			Road string `json:"road"`
		} `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Header.Add need to be put before content writing?
	w.Header().Add("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
