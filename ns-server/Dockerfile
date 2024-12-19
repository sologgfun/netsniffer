# 使用 ubuntu 作为基础镜像
FROM ubuntu:20.04

# 设置时区为上海
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 安装 MySQL 及其他依赖
RUN apt-get update && \
    apt-get install -y \
    mysql-server \
    curl \
    gnupg \
    lsb-release \
    && rm -rf /var/lib/apt/lists/*

# 将所需的脚本文件复制到容器中
COPY ./mysql/start.sh /mysql/start.sh
COPY ./mysql/schema.sql /mysql/schema.sql
COPY /ns-server /ns-server
COPY /ns-ctl /ns-ctl

# 确保脚本可执行
RUN chmod +x /mysql/start.sh

# 设置容器启动时运行的命令
CMD ["sh", "/mysql/start.sh"]