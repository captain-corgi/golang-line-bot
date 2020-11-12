run:
	go run cmd/echo_bot/server.go

tidy:
	go mod tidy
	go mod vendor