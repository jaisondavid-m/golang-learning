package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("New request received")
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Host:", r.Host)
	fmt.Println("RemoteAddr (IP & Port):", r.RemoteAddr)
	
	clientIP := r.RemoteAddr
	fmt.Println("Client IP:", clientIP)
	
	fmt.Println("User-Agent:", r.UserAgent())
	fmt.Println("Referer:", r.Referer())
	fmt.Println("Accept-Language:", r.Header.Get("Accept-Language"))
	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))

	queryParams := r.URL.Query()
	fmt.Println("Query Parameters:")
	for key, values := range queryParams {
		fmt.Printf("  %s: %v\n", key, values)
	}

	fmt.Fprintln(w, "Request completed successfully")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on :8000")
	
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}