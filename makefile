.PHONY: proto clean run w-players ms-players ms-teams ms-users ms-tournaments ms-matches

proto:
	protoc --go_out=.. --go-grpc_out=.. proto/*.proto
	
clean:
	del /s /q *.exe || exit 0
	powershell -Command "Get-ChildItem -Recurse -Directory -Filter logs | Remove-Item -Recurse -Force"

# API GATEWAY
run:
	go build -o cmd/api_gateway/api-gateway.exe cmd/api_gateway/main.go
	./cmd/api_gateway/api-gateway.exe
	

# WORKERS
worker:
	go build -o cmd/workers/worker.exe cmd/workers/worker.go
	./cmd/workers/worker.exe


# MICROSERVICES
ms-players:
	go build -o cmd/microservices/players/microservice-players.exe cmd/microservices/players/main.go
	./cmd/microservices/players/microservice-players.exe

ms-teams:
	go build -o cmd/microservices/teams/microservice-teams.exe cmd/microservices/teams/main.go
	./cmd/microservices/teams/microservice-teams.exe

ms-users:
	go build -o cmd/microservices/users/microservice-users.exe cmd/microservices/users/main.go
	./cmd/microservices/users/microservice-users.exe

ms-tournaments:
	go build -o cmd/microservices/tournaments/microservice-tournaments.exe cmd/microservices/tournaments/main.go
	./cmd/microservices/tournaments/microservice-tournaments.exe

ms-matches:
	go build -o cmd/microservices/matches/microservice-matches.exe cmd/microservices/matches/main.go
	./cmd/microservices/matches/microservice-matches.exe