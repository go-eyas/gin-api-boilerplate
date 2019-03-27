FROM alpine

COPY api /opt/app/
COPY config.toml /opt/app/

ENV API_DEBUG=false

WORKDIR /opt/app

CMD [ "api", "server" ]

