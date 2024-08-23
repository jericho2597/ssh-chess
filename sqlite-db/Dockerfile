# sqlite-db/Dockerfile
FROM alpine:latest

RUN apk add --no-cache sqlite

COPY init-db.sh /init-db.sh
RUN chmod +x /init-db.sh

CMD ["/init-db.sh"]