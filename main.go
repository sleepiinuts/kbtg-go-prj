package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	// GET: ping
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Add("Content-type", "application/json")
		w.Write([]byte("pong"))
	})

	//

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

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

		if err := json.NewEncoder(w).Encode(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-type", "application/json")
		// w.Write([]byte(fmt.Sprintf("Hello, %v", req)))
	})
	http.ListenAndServe(":8080", nil)

}
