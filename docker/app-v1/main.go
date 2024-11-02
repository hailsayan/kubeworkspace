package main

import (
	"fmt"
	"net/http"
	"os"
)

func H(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello from %s\n", os.Getenv("HOSTNAME"))
}

func main() {
	http.HandleFunc("/", H)

	http.ListenAndServe(":8080", nil)
}
