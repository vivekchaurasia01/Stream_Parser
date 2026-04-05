APP=Stream_Parser

build:
	go build -o $(APP) main.go

run: build
	./$(APP) $(FILE)