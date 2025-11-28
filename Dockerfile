FROM golang:1.24-alpine AS deps

WORKDIR /app

COPY container_src/go.mod ./
RUN go mod download

FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY --from=deps /go/pkg /go/pkg
COPY container_src/go.mod ./
COPY container_src/*.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /server

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /server /server

EXPOSE 8080


CMD ["/server"]
