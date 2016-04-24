FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf \
    /var/cache/apk/*

ADD bin/solder-cli /usr/bin/
ENTRYPOINT ["/usr/bin/solder-cli"]
CMD ["help"]
