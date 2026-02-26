ARG GO_VERSION=1.25
# arg version

FROM golang:${GO_VERSION}-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

#runtime
FROM ubuntu:25.10

# user
# cert

WORKDIR /app

COPY --from=builder app .

EXPOSE 8080

CMD ["./app"]
