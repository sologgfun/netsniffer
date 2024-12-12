#!/bin/bash

# 启动 MySQL 服务
mysqld &

while ! mysqladmin ping --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 1
done

echo '1.开始导入数据....'
mysql < /mysql/schema.sql
echo '2.导入数据完毕....'

mysql -u root <<EOF
ALTER USER 'root'@'localhost' IDENTIFIED WITH 'mysql_native_password' BY 'rootpwd';
FLUSH PRIVILEGES;
EOF

# 设置前端文件目录
export FRONTEND_DIR=/frontend/dist

# 启动 kt-npd-server
/kt-npd-server