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
# Install dependencies, imagemagick, go, and cleanup
RUN apt-get update && apt-get install -y \
    gcc pkg-config \
&& go get -u \
&& go build -o natricon \
&& rm -rf /var/lib/apt/lists/*

# Only copy binary
FROM base

USER root
WORKDIR /root

COPY --from=builder /usr/src/app/natricon /usr/bin/natricon

ENV GIN_MODE="release"

# run main.go
CMD ["natricon", "-host=0.0.0.0", "-port=5555", "-logtostderr"]
