run:
	go run github.com/atkinsonbg/go-gmux-proper-unit-testing

up:
	docker-compose down
	COMPOSE_DOCKER_CLI_BUILD=1 docker-compose up --build --detach
	docker-compose logs -f

down:
	rm -rf go-gmux-proper-unit-testing-api
	docker-compose down

docker:
	docker build -t github.com/atkinsonbg/go-gmux-proper-unit-testing/api:latest .

dockerrun:
	docker run -it github.com/atkinsonbg/go-gmux-proper-unit-testing/api:latest

dockertest:
	docker build -f Dockerfile.test -t atkinsonbg/go-postgres-test:local .

dockertestrun:
	docker run -it github.com/atkinsonbg/go-gmux-proper-unit-testing/tests:latest

testlocal:	dockertest
	docker run -v ${PWD}/cover.out:/testdir/cover.out -e GIT_URL='' atkinsonbg/go-postgres-test:local

testremote:
	docker run -v ${PWD}/cover.out:/testdir/cover.out -e GIT_URL='https://github.com/atkinsonbg/go-gmux-unit-testing.git' atkinsonbg/go-postgres-test:latest