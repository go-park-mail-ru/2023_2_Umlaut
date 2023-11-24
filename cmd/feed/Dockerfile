FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o  ./cmd/feed/main.go

RUN chmod +x ./auth

CMD ["./feed"]