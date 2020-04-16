FROM golang:1.14.2-buster

WORKDIR /go/src/app
# add source code
ADD . .
# add assets
ADD assets assets
# Dependencies and build
RUN go get -u
RUN go build
# run main.go
CMD ["./natricon", "-host=0.0.0.0", "-port=5555", "-logtostderr"]