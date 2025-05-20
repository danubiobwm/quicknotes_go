package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		message := string(body)
		fmt.Println("Received message:", message)

		if strings.Contains(message, "sair") {
			fmt.Println("Client requested disconnect (HTTP context)")
			fmt.Fprintf(w, "Server acknowledged 'sair'. Disconnecting from HTTP perspective.")
			return
		}

		fmt.Fprintf(w, "Your message was received successfully: %s", message)
	})

	fmt.Println("HTTP server listening on :5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
