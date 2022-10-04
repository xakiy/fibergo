build:
	go build -o server.bin main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run