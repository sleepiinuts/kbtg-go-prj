package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sleepiinuts/kbtg-go-prj/internal/ping"
)

func main() {
	fmt.Println("hello world")

	// GET: ping
	http.HandleFunc("/ping", ping.PingHandler)

	//

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/hello2", HelloHandler)
	http.ListenAndServe(":8080", nil)

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
