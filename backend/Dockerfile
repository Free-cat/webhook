FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /cmd/rest_server/main cmd/rest_server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /cmd/rest_server/main /main
EXPOSE 8080
CMD ["/main"]

