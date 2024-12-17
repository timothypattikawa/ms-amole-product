FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY application-*.yml .

EXPOSE 9092 9092
CMD [ "/app/main" ]