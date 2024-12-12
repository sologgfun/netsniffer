package main

import (
	"fmt"
	"kyanos-server/httpHandler"
	"net/http"
)

func main() {
	httpHandler.InitDB()
	http.HandleFunc("/api/records", httpHandler.GetConnectionRecords)
	http.HandleFunc("/api/save", httpHandler.PostConnectionRecord)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
