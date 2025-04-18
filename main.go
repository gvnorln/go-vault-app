package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Vault! User: %s, Pass: %s", user, pass)
	})

	http.ListenAndServe(":8080", nil)
}
