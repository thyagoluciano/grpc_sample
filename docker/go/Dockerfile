FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o main /app/cmd/server

EXPOSE 8080

CMD ["/app/main","-tls"]