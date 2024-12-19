package main

import (
	"fmt"
	"kyanos/ns-ctl/gui"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 定义命令行参数
	// req := flag.String("req", "", "Request parameter")
	// flag.Parse()
	// // 检查是否提供了 -req 参数
	// if *req == "" {
	// 	fmt.Println("Usage: netsniffer -req <request>")
	// 	os.Exit(1)
	// }
	// // 打印参数
	// fmt.Printf("Request parameter: %s\n", *req)
	InitDB()
}

var db *gorm.DB

func InitDB() {
	fmt.Printf("start to connect database\n")
	var err error
	dsn := "root:rootpwd@tcp(127.0.0.1:3306)/kyanos_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("start to print records\n")
	gui.PrintRecords(db)
}
