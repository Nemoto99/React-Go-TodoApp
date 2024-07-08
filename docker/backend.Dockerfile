FROM golang:1.21


RUN go install github.com/cosmtrek/air@v1.29.0

WORKDIR /go/src/app

CMD ["air"]