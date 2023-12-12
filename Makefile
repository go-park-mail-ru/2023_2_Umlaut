# path to docker compose file
DCOMPOSE:=docker-compose.yaml

# improve build time
DOCKER_BUILD_KIT:=COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1

down:
	docker compose -f ${DCOMPOSE} down --remove-orphans

build:
	${DOCKER_BUILD_KIT} docker compose build

run:
	docker compose up --build

test:
	go test -coverpkg=./... -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v "mocks\|cmd\|configs\|db\|docs\|model\|monitoring\|static\|utils\|dto.go\|service.go\|handler.go\|repository.go\|server\|.pb.go\|.pb\|proto\|middleware.go" > cover.out
	go tool cover -func=cover.out

swag:
	swag init -g cmd/app/main.go