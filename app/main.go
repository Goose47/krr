package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/sleep", sleep)

	http.ListenAndServe(":8000", nil)
}

// ping responds with "pong" to every request
func ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

// sleep sleeps for amount of seconds specified in sleep query param if present
// and responds with amount of seconds slept
func sleep(w http.ResponseWriter, req *http.Request) {
	t := req.URL.Query().Get("sleep")

	seconds := 0
	if t != "" {
		var err error
		seconds, err = strconv.Atoi(t)
		if err != nil {
			w.Write([]byte("Bad parameter"))
			return
		}
		time.Sleep(time.Duration(seconds) * time.Second)
	}

	w.Write([]byte(fmt.Sprintf("slept for %ds", seconds)))
}
