FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bin/api ./cmd/api
RUN go build -o /bin/worker ./cmd/worker

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache docker-cli
RUN mkdir -p /app/temp

COPY --from=builder /bin/api /bin/api
COPY --from=builder /bin/worker /bin/worker

CMD ["/bin/api"]