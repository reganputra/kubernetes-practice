FROM alpine:latest

ADD docker-entrypoint.sh /

RUN chmod +x /docker-entrypoint.sh && apk update && apk add openssl

ENTRYPOINT ["/docker-entrypoint.sh"]
