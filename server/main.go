package main

import (
	"fmt"
	"kyanos-server/httpHandler"
	"net/http"
)

func main() {
	httpHandler.InitDB()
	http.HandleFunc("/api/records", httpHandler.GetConnectionRecords)
	http.HandleFunc("/save", httpHandler.PostConnectionRecord)

	// 添加静态文件服务
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
