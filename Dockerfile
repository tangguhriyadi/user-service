FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /user_service

FROM alpine:latest


COPY --from=builder /user_service /user_service
CMD ["/user_service"]
