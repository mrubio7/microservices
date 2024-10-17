.PHONY: proto clean build run build-workers

proto:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. players.proto

build:
	go build -o cmd/api_gateway/api-gateway.exe cmd/api_gateway/api-gateway.go

run: 
	./cmd/api_gateway/api-gateway.exe
	
clean:
	del /s /q *.exe
	if exist logs rmdir /s /q logs & mkdir logs

build-workers:
	go build -o cmd/workers/players/worker-players.exe cmd/workers/players/worker-players.go

build-microservices:
	go build -o cmd/microservices/players/microservice-players.exe cmd/microservices/players/main.go

run-ms-players:
	./cmd/microservices/players/microservice-players.exe