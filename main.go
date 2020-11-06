package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	remoteAddr := r.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	fmt.Fprintf(w, "%s\n", remoteAddr)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Fprintf(os.Stderr, "error: set $PORT to something\n")
		os.Exit(1)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
