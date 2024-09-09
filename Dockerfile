FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/app

FROM alpine:latest AS production

WORKDIR /app

COPY --from=builder /go/bin/app /app/app.bin

CMD ["/app/app.bin"]


