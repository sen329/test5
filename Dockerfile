FROM golang:1.16.4
RUN mkdir /gm-tool-backend

ADD . /gm-tool-backend

WORKDIR /gm-tool-backend

RUN go build -o main

CMD ["/gm-tool-backend/main"]