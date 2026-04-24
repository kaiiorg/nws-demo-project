build: tidy
	mkdir -p ./bin
	go build -o ./bin ./cmd/forecast-api

tidy:
	go mod tidy

fmt:
	gofmt -w -s .

run:
	./bin/forecast-api

test: tidy
	go test ./...

call:
	curl -d '{"latitude": 36.7158451, "longitude": -91.8739187}' http://localhost:8080/api/v1/forecast

call-invalid:
	curl -d '{"latitude": 1000.0, "longitude": -91.8739187}' http://localhost:8080/api/v1/forecast

call-nws-invalid:
	curl -d '{"latitude": -36.7158451, "longitude": -91.8739187}' http://localhost:8080/api/v1/forecast
