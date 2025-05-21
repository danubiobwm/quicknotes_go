package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	fmt.Print("Servidor rodando na porta 5000\n")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	http.ListenAndServe(":5000", mux)
}
