FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o  ./cmd/auth/main.go

RUN chmod +x ./auth

CMD ["./auth"]