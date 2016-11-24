package main

import (
	"fmt"
	"net/http"
	"time"
)

// LaasHandler returns OK
func LaasHandler(w http.ResponseWriter, r *http.Request) {
	duration, err := time.ParseDuration(r.FormValue("duration"))
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		time.Sleep(duration)
		fmt.Fprint(w, "OK")
	}
}

func main() {
	http.HandleFunc("/latency", LaasHandler)
	http.ListenAndServe(":8080", nil)
}
