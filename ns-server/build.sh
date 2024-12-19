# 对ns-ctl和ns-server目录按顺序进行make image操作
# 1. ns-ctl
# 2. ns-server

cd ../ns-ctl
make image
cd ../ns-server
make image