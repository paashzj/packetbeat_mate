FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o packetbeat_mate .


FROM ttbb/packetbeat:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/packetbeat/mate

COPY --from=build /opt/sh/compile/pkg/packetbeat_mate /opt/sh/packetbeat/mate/packetbeat_mate

WORKDIR /opt/sh/packetbeat

CMD ["/usr/local/bin/dumb-init", "bash", "-vx", "/opt/sh/packetbeat/mate/scripts/start.sh"]