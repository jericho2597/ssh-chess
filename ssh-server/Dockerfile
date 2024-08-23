FROM golang:1.20.5-alpine
WORKDIR /app

COPY ./certs/ssh_host_key /app/ssh_host_key

COPY . .
RUN go mod tidy

RUN go build -o ssh-server ./cmd/.
RUN chmod +x ssh-server

EXPOSE 8080

CMD ["./ssh-server"]
