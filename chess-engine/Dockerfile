FROM golang:1.20.5-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o chessengine ./server/server.go
RUN chmod +x chessengine
CMD ["./chessengine"]