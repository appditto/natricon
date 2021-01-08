FROM appditto/imagemagick_go:latest as base

FROM base as builder

WORKDIR /root

# add source code
ADD . .
# add assets
ADD assets assets
# Dependencies and buil
# Install dependencies, imagemagick, go, and cleanup
RUN apt-get update && apt-get install -y \
    gcc pkg-config \
&& go get \
&& echo "Building natricon" \
&& go build -o natricon \
&& rm -rf /var/lib/apt/lists/*

# Copy binary to imagemagick image
FROM appditto/imagemagick:latest

USER root
WORKDIR /root

COPY --from=builder /root/natricon /usr/bin/natricon

ENV GIN_MODE="release"

# run main.go
CMD ["natricon", "-host=0.0.0.0", "-port=5555", "-logtostderr"]
