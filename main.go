package main

import (
	"fmt"
	"net/http"

	"github.com/sleepiinuts/kbtg-go-prj/internal/ping"
)

func main() {
	fmt.Println("hello world")

	// GET: ping
	http.HandleFunc("/ping", ping.PingHandler)

	http.HandleFunc("/hello", ping.HelloHandler)
	http.ListenAndServe(":8080", nil)

}
