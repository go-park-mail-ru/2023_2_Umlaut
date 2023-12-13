# path to docker compose file
DCOMPOSE:=docker-compose.yml

# improve build time
DOCKER_BUILD_KIT:=COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1

all: build up

down:
	docker compose -f ${DCOMPOSE} down --remove-orphans

build:
	${DOCKER_BUILD_KIT} docker compose build

up:
	docker compose up -d --remove-orphans

# Vendoring is useful for local debugging since you don't have to
# reinstall all packages again and again in docker
mod:
	go mod tidy -compat=1.21 && go mod vendor && go install ./...

tests:
	go test -coverpkg=./... -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v "mocks\|cmd\|configs\|db\|docs\|model\|monitoring\|static\|utils\|dto.go\|service.go\|handler.go\|repository.go\|server\|.pb.go\|.pb\|proto\|middleware.go" > cover.out
	go tool cover -func=cover.out

mock:
	mockgen -source=pkg/repository/repository.go -destination=pkg/repository/mocks/mock.go \
	&& mockgen -source=pkg/service/service.go -destination=pkg/service/mocks/mock.go

swag:
	swag init -g cmd/app/main.go

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run