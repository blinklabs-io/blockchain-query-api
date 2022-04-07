FROM golang:1.17 AS build

COPY . /code
WORKDIR /code
RUN make build

FROM ubuntu:bionic
COPY --from=build /code/blockchain-query-api /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/blockchain-query-api"]
