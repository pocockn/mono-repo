test:
	@GO111MODULE=on go test ./... -cover

down:
	@echo "Taking down common services. \n"
	docker-compose down

run: vendor
	@echo "Starting up the services. \n"
	docker-compose pull
	docker-compose up -d --remove-orphans

vendor:
	@go mod vendor
