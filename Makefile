include .env
export

run:
	go run main.go

buildrun:
	go build -o app main.go
	./app