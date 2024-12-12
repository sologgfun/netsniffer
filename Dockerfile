# 使用 ubuntu 作为基础镜像
FROM ubuntu:20.04

# 设置时区为上海
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY /kyanos /kyanos
# 增加权限
RUN chmod 777 /kyanos
