OK_RESULT= echo "Done😌!"
CMD_ROOT_FILE_PATH=./main.go

build: ## Build application
	@echo "Building application..."
	@go build -o tmp/andy $(CMD_ROOT_FILE_PATH) && ${OK_RESULT}

dev: ## Dev mode
	@air -c .air.toml

# dev: ## Run build during any file changing (dev live mode)
# 	@echo "Restarting..."
# 	ulimit -n 1000
# 	reflex --start-service -r '\.go$$' make build

help: # Show Makefile commands
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'
