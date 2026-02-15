APP_NAME=todo-app
PORT=8080

.PHONY: run build test docker-build docker-run clean

run:
	go run .

build:
	go build -o $(APP_NAME)

test:
	go test ./...

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p $(PORT):$(PORT) $(APP_NAME)

clean:
	rm -f $(APP_NAME)
