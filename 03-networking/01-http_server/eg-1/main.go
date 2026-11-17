package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter,r *http.Request) {

	ctx := r.Context()

	fmt.Println("Request Recived")

	select {
	case <-time.After(5*time.Second):
		fmt.Println(r, "Request completed successfully")
	case <-ctx.Done():
		fmt.Println("Request cancelled:",ctx.Err())
		return
	}
}

func main() {

	http.HandleFunc("/",handler)

	fmt.Println("Server running on :8000")
	
	http.ListenAndServe(":8000",nil)
}