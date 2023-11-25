FROM golang:latest

RUN go version
ENV GOPATH=/

RUN apt-get update
RUN apt-get -y install postgresql-client

COPY ./ ./

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o admin ./cmd/admin/main.go

CMD ["./admin"]