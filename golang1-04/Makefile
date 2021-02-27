hello:
	echo "Hello"

build:
	go build -o bin/main *.go

run:
	go run .


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm *.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 *.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 *.go

test:
	go test

all: hello build