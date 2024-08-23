# chess-api/Dockerfile
FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go build -o chess-api .

CMD ["sh", "-c", "while :; do ./chess-api; sleep 86400; done"]