generate:
	docker run --rm -v ${CURDIR}:${CURDIR} -w ${CURDIR} znly/protoc --go_out=plugins=grpc:./server -I ${CURDIR}/proto producer.proto

install-ghz:
	brew install ghz

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

kafka-start:
	clear
	docker-compose -f "deployments/docker-compose.yml" up --build

kafka-stop:
	docker-compose -f "deployments/docker-compose.yml" down --remove-orphans
	rm -rf ./deployments/prometheus/data

docker-clear:
	docker stop $(shell docker ps -a -q) || docker rm $(shell docker ps -a -q)
