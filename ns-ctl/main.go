package main

import (
	"flag"
	"kyanos/ns-ctl/gui"
	"os"
	"runtime/debug"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"k8s.io/klog/v2"
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
	defer func() {
		// 检查是否有 panic 发生
		if r := recover(); r != nil {
			// 捕获 panic，并输出日志
			klog.Errorf("Caught panic: %v\nStack trace: %s", r, debug.Stack())
		}
	}()
	InitDB()
}

func init() {
	ekcsLogFile, err := os.OpenFile("/tmp/ns-ctl.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		klog.Errorf("open log file error: %s", err)
	}
	klog.InitFlags(nil)
	err = flag.Set("logtostderr", "false")
	if err != nil {
		klog.Errorf("flag set error: %s", err)
		os.Exit(1)
	}
	err = flag.Set("alsologtostderr", "false")
	if err != nil {
		klog.Errorf("flag set error: %s", err)
		os.Exit(1)
	}
	klog.SetOutputBySeverity("INFO", ekcsLogFile)
}

var db *gorm.DB

func InitDB() {
	klog.Infof("start to connect database")
	var err error
	dsn := "root:rootpwd@tcp(127.0.0.1:3306)/kyanos_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	klog.Infof("start to print records")
	gui.PrintRecords(db)
}
