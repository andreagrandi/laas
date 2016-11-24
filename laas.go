package main

import (
	"fmt"
	"net/http"
	"time"
)

// LaasHandler returns OK
func LaasHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(500 * time.Millisecond)
	fmt.Fprint(w, "OK")
}

func main() {
	http.HandleFunc("/latency", LaasHandler)
	http.ListenAndServe(":8080", nil)
}
