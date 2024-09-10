FROM node:14.17-alpine AS node-env
WORKDIR /node
# copy only package*.json to cache go dependencies
COPY ./ssh-server/web/package*.json ./web/
RUN cd web && npm install --production

COPY ./ssh-server/web ./web

FROM golang:1.20.5-alpine
WORKDIR /app

ARG SSH_CHESS_SERVER_KEY_PATH
ARG SSH_CHESS_CLIENT_KEY_PATH

# put the configured ssh host key in the container at expected path
COPY ${SSH_CHESS_SERVER_KEY_PATH} /certs/ssh_host_key
COPY ${SSH_CHESS_CLIENT_KEY_PATH} /certs/ssh_client_key

# copy only go.mod and go.sum to cache go dependencies
COPY ./ssh-server/go.mod ./ssh-server/go.sum ./
RUN go mod download

COPY ./ssh-server/. .

RUN go build -o servers ./cmd/.
RUN chmod +x servers

COPY --from=node-env /node/web ./web

EXPOSE 80
EXPOSE 22

CMD ["./servers"]
