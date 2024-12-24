set @@global.sql_mode ='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

create database `kyanos_server` default character set utf8mb4 collate utf8mb4_general_ci;

use kyanos_server;

CREATE TABLE IF NOT EXISTS annotated_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    local_port INT UNSIGNED DEFAULT 0,
    remote_port INT UNSIGNED DEFAULT 0,
    remote_addr VARCHAR(45) DEFAULT '',
    local_addr VARCHAR(45) DEFAULT '',
    pid INT UNSIGNED DEFAULT 0,
    pid_str VARCHAR(255) DEFAULT '',
    protocol INT UNSIGNED DEFAULT 0,
    side TINYINT DEFAULT 0,
    stream_id INT DEFAULT 0,
    is_ssl BOOLEAN DEFAULT FALSE,
    req_str TEXT,
    resp_str TEXT,
    req_plain_text_size INT DEFAULT 0,
    resp_plain_text_size INT DEFAULT 0,
    req_size INT DEFAULT 0,
    resp_size INT DEFAULT 0,
    total_duration FLOAT DEFAULT 0.0,
    black_box_duration FLOAT DEFAULT 0.0,
    copy_to_socket_buffer_duration FLOAT DEFAULT 0.0,
    read_from_socket_buffer_duration FLOAT DEFAULT 0.0,
    start_ts BIGINT UNSIGNED DEFAULT 0,
    end_ts BIGINT UNSIGNED DEFAULT 0,
    req_syscall_event_details_json TEXT,
    resp_syscall_event_details_json TEXT,
    req_nic_event_details_json TEXT,
    resp_nic_event_details_json TEXT
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;