FROM appditto/libvips:latest as builder

ARG TARGETARCH

WORKDIR /root

# Update GO
# Install GOLang
RUN apt-get update && apt-get install -y wget \
  && wget -c https://dl.google.com/go/go1.22.linux-${TARGETARCH}.tar.gz -O - | tar -xz -C /usr/local \
  && rm -rf go1.22.linux-${TARGETARCH}.tar.gz \
  && rm -rf /var/lib/apt/lists/*
ENV PATH="/usr/local/go/bin:${PATH}"

# add source code
ADD . .
# add assets
ADD assets assets
# Build
RUN  go build -o /root/natricon

FROM ubuntu:22.04

ADD assets assets
COPY --from=builder /root/natricon /usr/local/bin/natricon

ENV MALLOC_ARENA_MAX=2
ENV GIN_MODE="release"

# Install libvips
RUN apt-get update && apt-get install -y \
    libvips-dev \
    && rm -rf /var/lib/apt/lists/*

# run main.go
CMD ["natricon", "-host=0.0.0.0", "-port=5555", "-logtostderr"]
