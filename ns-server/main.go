package main

import (
	"fmt"
	"kyanos-server/httpHandler"
	"net/http"
	"os"
)

func main() {
	currentPid := os.Getpid()
	mysqlPid, err := httpHandler.GetMySQLPid()
	if err != nil {
		fmt.Println("Failed to get MySQL process PID")
		return
	}
	// print pid
	fmt.Printf("current pid: %d, mysql pid: %d\n", currentPid, mysqlPid)
	httpHandler.InitDB()
	http.HandleFunc("/api/records", httpHandler.GetConnectionRecords)
	http.HandleFunc("/api/save", httpHandler.PostConnectionRecord)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
