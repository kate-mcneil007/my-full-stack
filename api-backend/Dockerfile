FROM alpine

COPY ./api-backend/api-backend.bin /

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN apk update && apk add ca-certificates curl && apk add bash && rm -rf /var/cache/apk/*

ENTRYPOINT ["/api-backend.bin"]