FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * ./

RUN go build -o /web-server

EXPOSE 8080

CMD [ "/web-server" ]