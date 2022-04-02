VERSION 0.6
FROM golang:1.18.0-alpine3.15
WORKDIR "/app"
RUN apk update && apk upgrade
RUN apk add --no-cache build-base rust cargo imagemagick

build:
    ENV MIX_ENV="prod"
    RUN mix do deps.get --only $MIX_ENV, deps.compile
    COPY config/config.exs config/$MIX_ENV.exs config/runtime.exs config/
    COPY assets assets
    COPY priv priv
    COPY rel rel
    COPY lib lib
    RUN mix do compile --warnings-as-errors, assets.clean, assets.deploy, release
    SAVE IMAGE --cache-hint
    SAVE ARTIFACT /app/_build/prod/rel/bazzile

docker:
    FROM alpine:3.14
    ENV LANG="en_US.UTF-8"
    ENV LANGUAGE="en_US:en"
    ENV LC_ALL="en_US.UTF-8"
    ENV DATABASE_URL="ecto://{user}:{password}@{host}/postgres"
    ENV CLUSTERING_DNS_QUERY="cp.{app_name}.hidora.com"
    RUN apk update && apk upgrade
    RUN apk add --no-cache ca-certificates imagemagick ncurses-libs && update-ca-certificates
    WORKDIR "/app"
    RUN chown nobody:nobody /app
    COPY --chown nobody:nobody +build/* . 
    USER nobody
    EXPOSE 4000
    CMD ["/app/bin/start"]
    ARG TAG="latest"
    SAVE IMAGE --push bazzile/app:$TAG
