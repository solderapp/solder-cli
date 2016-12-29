FROM alpine:edge
MAINTAINER Thomas Boerger <thomas@webhippie.de>

RUN apk update && \
  apk add \
    ca-certificates \
    bash && \
  rm -rf \
    /var/cache/apk/* && \
  addgroup \
    -g 1000 \
    kleister && \
  adduser -D \
    -h /home/kleister \
    -s /bin/bash \
    -G kleister \
    -u 1000 \
    kleister

COPY kleister-cli /usr/bin/

USER kleister
ENTRYPOINT ["/usr/bin/kleister-cli"]
CMD ["help"]
