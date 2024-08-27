FROM golang:1.22-alpine AS builder

COPY . /github.com/marktsarkov/notes-app-kode/
WORKDIR /github.com/marktsarkov/notes-app-kode/

RUN go mod download
RUN go build -o ./bin/server ./cmd/server

COPY go.mod go.sum ./

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/marktsarkov/notes-app-kode/bin/server .

CMD ["./server"]

