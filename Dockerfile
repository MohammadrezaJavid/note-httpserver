ARG     CODE_VERSION=1.21rc2-alpine3.18

FROM    golang:${CODE_VERSION}

ENV     HOME=/httpServers

WORKDIR ${HOME}

RUN mkdir -p ${HOME} &&\
    mkdir -p ${HOME}/txt &&\
    mkdir -p ${HOME}/html &&\
    mkdir -p ${HOME}/httpServer

COPY ./go.mod           ${HOME}
COPY ./*.go             ${HOME}
COPY ./httpServer/*.go  ${HOME}/httpServer
COPY ./html/*.html      ${HOME}/html

RUN go build -o httpserver main.go 

EXPOSE 8080

CMD [ "./httpserver" ]