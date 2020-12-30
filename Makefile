include .env
export

run:
	go run github.com/atkinsonbg/go-gmux-db-testcontainers

up:
	docker-compose down
	COMPOSE_DOCKER_CLI_BUILD=1 docker-compose up --build --detach
	docker-compose logs -f

down:
	rm -rf go-gmux-proper-unit-testing-api
	docker-compose down

docker:
	docker build -t github.com/atkinsonbg/go-gmux-db-testcontainers:latest .

dockerrun:
	docker run -it github.com/atkinsonbg/go-gmux-db-testcontainers:latest

test:
	go test -v ./... -coverpkg ./...