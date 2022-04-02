VERSION 0.6
FROM golang:1.17.8-alpine3.15
WORKDIR "/app"
RUN apk update && apk upgrade
RUN apk add --no-cache build-base
RUN go install github.com/swaggo/swag/cmd/swag@latest

build:
    COPY docs docs
    COPY ent ent
    COPY server server
    COPY go.mod go.sum main.go .
    RUN ls -la
	RUN go run github.com/swaggo/swag/cmd/swag@latest init
    RUN go build
    SAVE ARTIFACT /app/calcio

docker:
    FROM alpine:3.15
    ENV LANG="en_US.UTF-8"
    ENV LANGUAGE="en_US:en"
    ENV LC_ALL="en_US.UTF-8"
    ENV DB_URL="host={host} port=5432 user={user} dbname=postgres password={password} sslmode=disable"
    RUN apk update && apk upgrade
    RUN apk add --no-cache ca-certificates && update-ca-certificates
    WORKDIR "/app"
    RUN chown nobody:nobody /app
    COPY --chown nobody:nobody +build/* . 
    RUN chmod +x calcio
    USER nobody:nobody
    EXPOSE 4000
    CMD ["./calcio"]    
    ARG TAG="latest"
    SAVE IMAGE --push bazzile/calcio:$TAG
