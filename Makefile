BINARY_NAME=challenge2019

statement1:
	go build -o ./${BINARY_NAME} ./cmd/main.go
	./${BINARY_NAME} cli statement1
statement2:
	go build -o ./${BINARY_NAME} ./cmd/main.go
	./${BINARY_NAME} cli statement2



