.PHONY: proto clean build run workers microservices w-players

proto:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. players.proto
	protoc --proto_path=proto --go_out=. --go-grpc_out=. teams.proto

build:
	go build -o cmd/api_gateway/api-gateway.exe cmd/api_gateway/main.go

run:
	go build -o cmd/api_gateway/api-gateway.exe cmd/api_gateway/main.go
	./cmd/api_gateway/api-gateway.exe
	
clean:
	del /s /q *.exe
	if exist logs rmdir /s /q logs & mkdir logs

workers:
	go build -o cmd/workers/players/worker-players.exe cmd/workers/players/worker-players.go

w-players: 
	./cmd/workers/players/worker-players.exe

microservices:
	go build -o cmd/microservices/players/microservice-players.exe cmd/microservices/players/main.go

ms-players:
	go build -o cmd/microservices/players/microservice-players.exe cmd/microservices/players/main.go
	./cmd/microservices/players/microservice-players.exe

ms-teams:
	go build -o cmd/microservices/teams/microservice-teams.exe cmd/microservices/teams/main.go
	./cmd/microservices/teams/microservice-teams.exe