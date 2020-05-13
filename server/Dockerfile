FROM appditto/imagemagick_go:latest as base

FROM base as builder

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app
RUN rm -rf /imagemagick-build

ENV CGO_CFLAGS_ALLOW="-Xpreprocessor"

# add source code
ADD . .
# add assets
ADD assets assets
# Dependencies and buil
RUN go get -u
RUN go build -o natricon

# Only copy binary
FROM base

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/natricon /usr/src/app/natricon

ENV GIN_MODE="release"

# run main.go
CMD ["./natricon", "-host=0.0.0.0", "-port=5555", "-logtostderr"]
