FROM golang:1.21-alpine AS go-builder
ARG arch=aarch64

ENV APPNAME=gotabitd

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.4.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.4.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep a8259ba852f1b68f2a5f8eb666a9c7f1680196562022f71bb361be1472a83cfd
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 324c1073cb988478d644861783ed5a7de21cfd090976ccc6b1de0559098fbbad

RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a

RUN set -eux; apk add --no-cache ca-certificates build-base;
RUN apk add git

WORKDIR /code

COPY . /code/

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build
RUN echo "Ensuring binary is statically linked ..." \
  && (file /code/build/$APPNAME | grep "statically linked")

FROM alpine:3.16

WORKDIR /chain

ENV APPNAME=gotabitd

COPY --from=go-builder /code/build/$APPNAME /usr/bin/$APPNAME

COPY ./scripts/docker/* /opt/

RUN chmod +x /opt/chain_*

# rest server
EXPOSE 1317
# grpc
EXPOSE 9090
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["sh", "-c", "/usr/bin/$APPNAME", "version"]

