OK_RESULT="Fine😌!"

build_app:
	@echo "Building application..."
	@go build -o build/andy ./src/.
	@echo ${OK_RESULT}

deploy_app:
	@echo "Moving application to /usr/local/bin/ folder..."
	@sudo cp ./build/andy /usr/local/bin/
	@echo ${OK_RESULT}

dev:
	@echo "Restarting..."
	ulimit -n 1000
	${GOPATH}/bin/reflex --start-service -r '\.go$$' make build_app
