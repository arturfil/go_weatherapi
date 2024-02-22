include .env

BINARY=weatherApi

build:
	@echo "Building backend"
	go build -o ${BINARY} cmd/main/main.go

run:
	@echo "Running binary"
	@env PORT=${PORT} SECRET_KEY=${SECRET_KEY} ./${BINARY} & 
	@echo "Backend running!"
	
stop:
	@echo "Stopping backend"
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "Stopped backend"
	
restart: stop build run

all: build run
