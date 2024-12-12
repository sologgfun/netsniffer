package httpHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

type Response struct {
	Message string `json:"message"`
}

type ConnectionRecord struct {
	ID                int       `json:"id"`
	ConnectionDesc    string    `json:"connection_desc"`
	Protocol          string    `json:"protocol"`
	TotalTimeMs       float64   `json:"total_time_ms"`
	RequestSize       int       `json:"request_size"`
	ResponseSize      int       `json:"response_size"`
	Process           string    `json:"process"`
	NetInternalTimeMs float64   `json:"net_internal_time_ms"`
	ReadSocketTimeMs  float64   `json:"read_socket_time_ms"`
	StartTime         time.Time `json:"start_time"`
	Request           string    `json:"request"`
	Response          string    `json:"response"`
}

func PostConnectionRecord(w http.ResponseWriter, r *http.Request) {
	var record ConnectionRecord
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
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
	var records []ConnectionRecord
	result := db.Find(&records)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}
