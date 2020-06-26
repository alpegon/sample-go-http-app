FROM golang:1.14.4-alpine3.12

ENV COLOUR blue
ENV LOAD_TIME 0
ENV RESPONSE_TIME 0
ENV PORT 8081
ENV VERSION 1.0.0

RUN mkdir /app
ADD *.go /app

WORKDIR /app

RUN apk add git && \ 
  go get github.com/gorilla/mux && \
  go build -o main . && \
  addgroup -S utils && \
  adduser -S -g utils utils

USER utils

CMD ["/app/main"]
