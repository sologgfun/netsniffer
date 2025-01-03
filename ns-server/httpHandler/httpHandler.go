package httpHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := "root:rootpwd@tcp(127.0.0.1:3306)/kyanos_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
func GetMySQLPid() (int, error) {
	cmd := exec.Command("pgrep", "mysqld")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	pidStr := strings.TrimSpace(string(output))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return 0, err
	}
	return pid, nil
}

func PostConnectionRecord(w http.ResponseWriter, r *http.Request) {
	var record AnnotatedRecord
	//打印r.Body的字符串
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.Create(&record).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Record inserted successfully")
}

func GetConnectionRecords(w http.ResponseWriter, r *http.Request) {
	var records []AnnotatedRecord
	result := db.Find(&records)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}
