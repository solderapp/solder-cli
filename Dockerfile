FROM alpine:edge

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

ARG VERSION
COPY dist/binaries/kleister-cli-$VERSION-linux-amd64 /usr/bin/

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>"
LABEL org.label-schema.version=$VERSION
LABEL org.label-schema.name="Kleister CLI"
LABEL org.label-schema.vendor="Thomas Boerger"
LABEL org.label-schema.schema-version="1.0"

USER kleister
ENTRYPOINT ["/usr/bin/kleister-cli"]
CMD ["help"]
