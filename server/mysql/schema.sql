create database `kyanos_server` default character set utf8mb4 collate utf8mb4_general_ci;

use kyanos_server;

-- 建表
CREATE TABLE IF NOT EXISTS connection_records (
        id INT AUTO_INCREMENT PRIMARY KEY,
        connection_desc VARCHAR(255) DEFAULT '',
        protocol VARCHAR(50) DEFAULT '',
        total_time_ms FLOAT DEFAULT 0.0,
        request_size INT DEFAULT 0,
        response_size INT DEFAULT 0,
        process VARCHAR(255) DEFAULT '',
        net_internal_time_ms FLOAT DEFAULT 0.0,
        read_socket_time_ms FLOAT DEFAULT 0.0,
        start_time DATETIME DEFAULT CURRENT_TIMESTAMP,
        request TEXT,
        response TEXT
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;