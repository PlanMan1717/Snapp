package main

import (
	"fmt"
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Go on macOS!\n"))

}

func handleSend(w http.ResponseWriter, r *http.Request) {
	recipientQuery := r.URL.Query()

	recipient := recipientQuery.Get("recipient")
	fmt.Println("To:", recipient)

	w.WriteHeader(http.StatusForbidden)
}

func handleUserAgent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	userAgent := r.UserAgent()
	_, err := fmt.Fprint(w, userAgent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("Hello, Go on macOS!")
	http.HandleFunc("/", handleGet)
	http.HandleFunc("/send", handleSend)
	http.HandleFunc("/whoami", handleUserAgent)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
