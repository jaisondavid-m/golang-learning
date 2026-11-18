package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Failed to read Body", http.StatusBadRequest)
		}

		responseBody := string(body)

		fmt.Printf("Recieved Data: ")
		fmt.Println(responseBody)
	default:
		http.Error(w,"This Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}

}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Server is listing on port :8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server Failed to start:", err)
	}

}
