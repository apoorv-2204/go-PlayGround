run_main:
	go run src/main/main.go
hello:
	echo "Hello"

build:
	go build -o bin/main main.go
	go build -o bin/dsa  DSA/*

# run:
# 	# go build src/main/main.go
# 	# ./go_playground
# 	go run src/main/main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
