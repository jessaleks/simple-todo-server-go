FROM golang:1.17-alpine AS BUILDER
RUN apk update && apk add --no-cache git

RUN mkdir $GOPATH/src/server

ADD * $GOPATH/src/server

WORKDIR $GOPATH/src/server

RUN go mod tidy
RUN go mod download
RUN mkdir /server
RUN go build -o /server/server

RUN go build -o /web-server


FROM alpine:latest
RUN mkdir /server
COPY --from=BUILDER /server/server /server/server
EXPOSE 8080
WORKDIR /server
CMD [ "/server/server" ]