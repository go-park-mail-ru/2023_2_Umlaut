# path to docker compose file
DCOMPOSE:=docker-compose.yml

# improve build time
DOCKER_BUILD_KIT:=COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1

down:
	docker compose -f ${DCOMPOSE} down --remove-orphans

build:
	${DOCKER_BUILD_KIT} docker compose build

up:
	docker compose up --build -d --remove-orphans

# Vendoring is useful for local debugging since you don't have to
# reinstall all packages again and again in docker
mod:
	go mod tidy -compat=1.21 && go mod vendor && go install ./...

tests:
	go test -coverpkg=./... -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v "mocks\|cmd\|configs\|db\|docs\|model\|monitoring\|constants\|utils\|service.go\|handler.go\|repository.go\|server.go|.pb.go\|.pb\|proto\|middleware.go" > cover.out
	go tool cover -func=cover.out

mock:
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock.go \
	&& mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go

swag:
	swag init -g cmd/app/main.go

lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run