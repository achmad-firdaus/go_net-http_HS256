# Base image Golang
FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o jwt-server

EXPOSE 8088

CMD ["./jwt-server"]
