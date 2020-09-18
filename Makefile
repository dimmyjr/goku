generate:
	docker run --rm -v ${CURDIR}:${CURDIR} -w ${CURDIR} znly/protoc --go_out=plugins=grpc:./server -I ${CURDIR}/proto producer.proto

install-ghz:
	brew install ghz

install-librdkafka:
	brew install librdkafka

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
