build: tidy
	go build -o ./bin ./cmd/forecast-api

tidy:
	go mod tidy

run:
	./bin/forecast-api
