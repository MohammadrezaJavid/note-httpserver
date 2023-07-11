FROM    golang:1.21rc2-alpine3.18 AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./

COPY *.go ./

COPY txt/ ./txt
COPY html/ ./html
COPY httpServer ./httpServer

RUN go mod tidy &&\
 CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o note-app 

CMD [ "./note-app" ]