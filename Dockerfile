FROM golang:1.14-alpine as builder

# install upx ca
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk add --no-cache upx ca-certificates tzdata

WORKDIR /usr/src/app
ENV GOPROXY=https://goproxy.cn

COPY ./go.mod /usr/src/app/
COPY ./go.sum /usr/src/app/
COPY ./basic/go.mod /usr/src/app/basic/
COPY ./basic/go.sum /usr/src/app/basic/


RUN go mod download

# build
COPY . /usr/src/app
RUN go build -ldflags "-s -w" -o server main.go && \
  mv server server_src && \
  upx --best server_src -o server -q


FROM alpine as runner

# fix timezone
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#   apk update && apk add tzdata && \
#   ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
#   echo "Asia/Shanghai" > /etc/timezone && \
#   rm -rf /var/cache/apk/* /tmp/* /var/tmp/* $HOME/.cache
RUN echo "Asia/Shanghai" > /etc/timezone

WORKDIR /opt/app

VOLUME /opt/app/runtime
EXPOSE 9000

ENV MK_DEBUG=false
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/src/app/api.toml /opt/app/
COPY --from=builder /usr/src/app/server /opt/app/

CMD ["/opt/app/server"]