FROM golang:1.20.5-alpine
WORKDIR /app

ARG SSH_CHESS_SERVER_KEY_PATH

# put the configured ssh host key in the container at expected path
COPY ${SSH_CHESS_SERVER_KEY_PATH} /certs/ssh_host_key

COPY ./ssh-server/. .
RUN go mod tidy

RUN go build -o ssh-server ./cmd/.
RUN chmod +x ssh-server

EXPOSE 8080

CMD ["./ssh-server"]
