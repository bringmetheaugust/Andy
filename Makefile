OK_RESULT="Done😌!"

help: # Show Makefile commands
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

build: ## Build application
	@echo "Building application..."
	@go build -o dist/andy ./cmd/andy.go
	@echo ${OK_RESULT}

run: ## Run application without build
	@echo "Be carefull!\nThis command does't work with the flags!\n"
	@go run ./cmd/andy.go $(filter-out $@, $(MAKECMDGOALS))

deploy: ## Deploy compiled application to your `/home/$USER/.local/bin` folder
	@echo "Moving application to /home/USER/.local/bin/ folder..."
	@sudo cp ./dist/andy /home/${USER}/.local/bin/
	@echo ${OK_RESULT}

dev: ## Run build during any file changing (dev live mode)
	@echo "Restarting..."
	ulimit -n 1000
	${GOPATH}/bin/reflex --start-service -r '\.go$$' make build
