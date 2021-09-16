FROM alpine:3
ENV LANG "en_US.UTF-8"
ENV LANGUAGE "en_US:en"
ENV LC_ALL "en_US.UTF-8"
RUN apk update && apk upgrade &&
COPY --chown=nobody:nobody calcio .
USER nobody:nobody
EXPOSE 4000
CMD ["./calcio"]
