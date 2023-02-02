# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.18-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /juncachain/
COPY go.sum /juncachain/
RUN cd /juncachain && go mod download

ADD . /juncachain
RUN cd /juncachain && go run build/ci.go install -static ./cmd/junca

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /juncachain/build/bin/junca /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["junca"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
