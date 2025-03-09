
FROM golang:1.24 AS builder
WORKDIR /stackgen
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .


FROM alpine:latest
WORKDIR /stackgen
COPY --from=builder /stackgen/main .
COPY --from=builder /stackgen/config config/
EXPOSE 8081
CMD ["./main"]
