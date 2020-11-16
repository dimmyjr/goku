clean:
	@rm -rf goku.out
	@rm -rf cover.out

build:
	@go env
	@go build -v

test:
	@go test -coverprofile=cover.out fmt ./...
	@go tool cover -html=cover.out

lint:
	golangci-lint run
	@#docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.32.2 golangci-lint run --timeout=10m

gocker:
	docker run -it --rm -v $(shell pwd):/usr/src/myapp -w /usr/src/myapp golang:1.13.15 bash
#	docker run -it --rm -v $(shell pwd):/usr/src/myapp -w /usr/src/myapp ubuntu:latest bash

generate:
	docker run --rm -v ${CURDIR}:${CURDIR} -w ${CURDIR} znly/protoc --go_out=plugins=grpc:./server -I ${CURDIR}/proto producer.proto

install-ghz:
	brew install ghz

install-librdkafka-macos:
	brew install librdkafka

install-librdkafka-linux:
	apt-get update && apt-get install build-essential pkg-config git librdkafka-dev -ygo

start-server:
	go run ./main.go -kafkaURL=localhost:9092 -topic=goku-sarama -provider=sarama -grpcPort=50051 -prometheusPort=8080

ghz: 
	ghz --insecure \
  --proto ./api/producer.proto \
  --call "server.Producer.publish" \
  --total 100000 \
  --concurrency 100 \
  --stream-interval=500ms \
  -d '{"message":"{{.RequestNumber}}"}' \
  0.0.0.0:50051

services-start:
	clear
	docker-compose -f "deployments/docker-compose.yml" up --build

services-stop:
	docker-compose -f "deployments/docker-compose.yml" down --remove-orphans
	rm -rf ./deployments/prometheus/data

docker-clear:
	docker stop $(shell docker ps -a -q) || docker rm $(shell docker ps -a -q)

kafka-consumer:
	docker run -it --network host confluentinc/cp-kafka:latest kafka-console-consumer --bootstrap-server localhost:9092 --from-beginning -timeout-ms 5000 --topic $(topic)
