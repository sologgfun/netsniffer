IMAGE_NAME ?= ccr.ccs.tencentyun.com/ktns/kt-npd-server
IMAGE_TAG ?= v0.1

image:
	rm -f kt-npd-server
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o kt-npd-server main.go
	sudo docker buildx build --platform linux/amd64 -t $(IMAGE_NAME):$(IMAGE_TAG) .
	rm -rf kt-npd-server