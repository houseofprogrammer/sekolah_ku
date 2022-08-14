FROM golang:1.18.3-alpine3.16 as builder
WORKDIR /app/go
COPY . /app/go
RUN go mod download
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o engine
#second stage
FROM alpine:3.16 as production
RUN addgroup -g 1001 -S gogogo && adduser -S gogogo -u 1001
WORKDIR /
RUN apk add --no-cache tzdata
COPY --from=builder /app/go .
#COPY .env .env
EXPOSE 8080

USER gogogo

ENTRYPOINT ["./engine"]