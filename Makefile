OK_RESULT="Done😌!"

build:
	@echo "Building application..."
	@go build -o build/andy ./cmd/andy.go
	@echo ${OK_RESULT}

deploy:
	@echo "Moving application to /usr/local/bin/ folder..."
	@sudo cp ./build/andy /usr/local/bin/
	@echo ${OK_RESULT}

dev:
	@echo "Restarting..."
	ulimit -n 1000
	${GOPATH}/bin/reflex --start-service -r '\.go$$' make build
