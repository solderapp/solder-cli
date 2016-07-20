FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf \
    /var/cache/apk/*

ADD bin/kleister-cli /usr/bin/
ENTRYPOINT ["/usr/bin/kleister-cli"]
CMD ["help"]
