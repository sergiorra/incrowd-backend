DC := docker-compose -p incrowd-backend -f docker/docker-compose.yml

run:
	go run cmd/incrowd-backend/main.go -conf config/config.local.json

# Stops and removes all Docker containers, networks, and volumes
docker-clean:
	$(DC) down --remove-orphans --volumes

# Cleans up containers and then starts up all the services defined in docker-compose
docker-up:
	make docker-clean
	$(DC) up --force-recreate -d

# Cleans up containers and then starts up only DB
docker-database:
	make docker-clean
	$(DC) up --force-recreate -d mongodb db-client

lint:
	golangci-lint run ./...