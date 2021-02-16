build_app:
	@go build -o andy ./src/index.go

deploy_app:
	@sudo cp ./andy /bin/
