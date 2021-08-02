FROM golang:1.16

WORKDIR /go/src/github.com/vitor-mariano/first-go-grpc-exercise

RUN apt update && apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

COPY . ./

RUN go mod tidy
RUN go build -o ./bin/server ./cmd/server/main.go
RUN go build -o ./bin/client ./cmd/client/main.go
