FROM golang:1.7.5-alpine
RUN mkdir -p /go/src/github.com/mikerap/hwserver
ADD . /go/src/github.com/mikerap/hwserver
RUN apk update && apk upgrade && apk add git
RUN go get github.com/gorilla/mux
RUN go install github.com/mikerap/hwserver
ENTRYPOINT /go/bin/hwserver
EXPOSE 8080
