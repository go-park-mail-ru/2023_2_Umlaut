build:
	docker compose build

run:
	docker compose up --build

test:
	go test -coverpkg=./... -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v "mocks\|cmd\|configs\|db\|docs\|model\|monitoring\|static\|utils\|dto.go\|service.go\|handler.go\|postgres.go\|server\|.pb.go\|.pb\|proto\|middleware.go" > cover.out
	go tool cover -func=cover.out

swag:
	swag init -g cmd/app/main.go