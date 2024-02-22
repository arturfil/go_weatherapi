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

get_miami_info:
	curl "http://localhost:${PORT}/api/get-weather?lat=25.761681&lon=-80.191788" | json_pp

restart: stop build run

all: build run
