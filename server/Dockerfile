FROM appditto/imagemagick_go:latest

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app
RUN rm -rf /imagemagick-build

ENV CGO_CFLAGS_ALLOW="-Xpreprocessor"
ENV GIN_MODE="release"

# add source code
ADD . .
# add assets
ADD assets assets
# Dependencies and buil
RUN go get -u
RUN go build -o natricon
# run main.go
CMD ["./natricon", "-host=0.0.0.0", "-port=5555", "-logtostderr"]
