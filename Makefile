build_app:
	@go build -o build/andy ./src/.

deploy_app:
	@sudo cp ./build/andy /usr/local/bin/

dev:
	@echo "Restarting..."
	ulimit -n 1000
	${GOPATH}/bin/reflex --start-service -r '\.go$$' make build_app
