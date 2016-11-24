package main

import (
	"fmt"
	"net/http"
)

// LaasHandler returns OK
func LaasHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	http.HandleFunc("/latency", LaasHandler)
	http.ListenAndServe(":8080", nil)
}