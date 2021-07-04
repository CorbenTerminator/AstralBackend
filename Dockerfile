FROM golang:latest

WORKDIR /app

COPY ./AstralBackend /app

RUN go build -o http_server

EXPOSE 8080

CMD ["./http_server"]