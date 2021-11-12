################################################################################
# build daemon stage
################################################################################
FROM golang:1.16.7-stretch AS build-env

ENV GO111MODULE=on

WORKDIR /appdebug
COPY . /appdebug

RUN set -ex \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
               build-essential

RUN git clone https://github.com/go-delve/delve.git /go/src/github.com/go-delve/delve && \
    cd /go/src/github.com/go-delve/delve && \
    git checkout v1.7.2 && \
    go install -ldflags "-X main.Build=$(git rev-parse HEAD) -extldflags \"-fno-PIC -static\"" -buildmode pie -tags 'osusergo netgo static_build' github.com/go-delve/delve/cmd/dlv

################################################################################
# merge
################################################################################
FROM amd64/debian:9-slim

RUN set -ex \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
               build-essential gcc git \
               ca-certificates openssl

# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
RUN update-ca-certificates

ENV PATH="/usr/local/go/bin:/usr/local/go:${PATH}"

COPY --from=build-env /appdebug /appdebug

WORKDIR /appdebug

COPY --from=build-env /go/bin/* /bin/
COPY --from=build-env /usr/local/go /usr/local/go
