TAG ?= dev
docker:
	docker build -t quay.io/ethanfrogers/golang-sample-service:$(TAG) .