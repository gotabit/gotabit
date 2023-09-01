FROM golang:1.20-alpine AS go-builder

ENV APPNAME=gotabitd

# See https://github.com/CosmWasm/wasmvm/releases
RUN ARCH=`uname -m`; echo ${ARCH}; \
  wget https://github.com/CosmWasm/wasmvm/releases/download/v1.2.1/libwasmvm_muslc.${ARCH}.a \
  -O /lib/libwasmvm_muslc.a; \
  # checksums
  wget https://github.com/CosmWasm/wasmvm/releases/download/v1.2.1/checksums.txt \
  -O /tmp/checksums.txt; \
  sha256sum /lib/libwasmvm_muslc.a | \
  grep $(cat /tmp/checksums.txt | grep ${ARCH} | cut -d ' ' -f 1)

RUN apk add --no-cache ca-certificates build-base git linux-headers;

WORKDIR /code

COPY . /code/

RUN BUILD_TAGS=muslc LINK_STATICALLY=true make build

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

