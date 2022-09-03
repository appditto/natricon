FROM appditto/libvips:latest

WORKDIR /root

# add source code
ADD . .
# add assets
ADD assets assets
# Dependencies and buil
# Install dependencies, libvips, go, and cleanup
RUN  go get \
    && go build -o natricon

ENV ENVIRONMENT "production"

ENV MALLOC_ARENA_MAX=2

ENV GIN_MODE="release"

# run main.go
CMD ["./natricon", "-host=0.0.0.0", "-port=5555", "-logtostderr"]
