build: tidy
	go build -o ./bin ./cmd/forecast-api

tidy:
	go mod tidy

fmt:
	gofmt -w -s .

run:
	./bin/forecast-api
