GOLANG_VERSION ?= 1.15.3
APP_NAME ?= crypto
APP_PORT ?= 3000

run:
	go run main.go

docker-run:
	docker run \
  		--name $(APP_NAME) \
  		-p $(APP_PORT):3000 \
  		-v $(PWD):/app -w /app \
  		--rm -d -t -i \
  		golang:$(GOLANG_VERSION) \
  		make run
