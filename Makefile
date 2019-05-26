.PHONY: all dev build env-up env-up-show env-down run clean

all: clean build env-up run

dev: build run

##### BUILD
build:
	@echo "Build ..."
	@go build
	@echo ""

##### ENV
env-up:
	@echo "Start enironment ..."
	@cd fixtures && docker-compose up --force-recreate -d
	@echo "Environment up"

env-up-show:
	@echo "Setup enironmet ..."
	@cd fixtures && docker-compose up
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd fixtures && docker-compose down
	@echo "Environment down"


##### RUN
run:
	@echo "Start app ..."
	@./kongyixueyuan

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/kongyixueyuan-* kongyixueyuan
	@docker rm -f -v `docker ps -a --no-trunc | grep "kongyixueyuan" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker rmi `docker images --no-trunc | grep "kongyixueyuan" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"

