build:
@go build -o bin/restapi-go cmd/main.go

test:
@ test -v

run: build

@./bin/restapi