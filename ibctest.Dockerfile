# To build the Desmos image, just run:
# > docker build -t desmos .
#
# Simple usage with a mounted data directory:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos desmos desmos init
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos desmos desmos start
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos --name desmos desmos
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it desmos bash
#
# To exit the bash, just execute
# > exit
FROM golang:1.15-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/desmos-labs/desmos

# Add source files
COPY . .

# Install Desmos, Relayer, remove packages
RUN make build-linux
RUN make get-relayer
RUN make build-relayer

# Generate ibc config
RUN make setup-ibctest


# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Install bash
RUN apk add --no-cache bash

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/build/desmos /usr/bin/desmos
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/.thirdparty/relayer/build/rly /usr/bin/rly

# Copy testnet cofig from build-env
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/build/ibc .
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/testconfig .

# Start two chains in background
RUN nohup desmos start --home ibc0 \
    --address tcp://0.0.0.0:26658 \
    --grpc.address 0.0.0.0:9090 \
    --p2p.laddr tcp://0.0.0.0:26656 \
    --rpc.laddr tcp://127.0.0.1:26657 \
    > /dev/null 2>&1 &
RUN nohup desmos start --home ibc1 \
    --address tcp://0.0.0.0:26668 \
    --grpc.address 0.0.0.0:9091 \
    --p2p.laddr tcp://0.0.0.0:26666 \
    --rpc.laddr tcp://127.0.0.1:26667 \
    > /dev/null 2>&1 &

# Setup relayer
RUN cd testconfig && setup.sh


EXPOSE 26656 26657 26658 9090 26668 9091 26666 26667
