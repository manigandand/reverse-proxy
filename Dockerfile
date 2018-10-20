FROM golang:1.10-alpine AS builder

MAINTAINER Manigandan Dharmalingam <manigandan.jeff@gmail.com>

COPY . /go/src/manigandand-golang-test/
WORKDIR /go/src/manigandand-golang-test/

RUN apk add --update bash make

RUN make build-server

# ------------

FROM alpine

COPY --from=builder /go/src/manigandand-golang-test/ /

RUN apk add --no-cache ca-certificates
ENV ENV developemnt
ENV PORT 8080
ENV API_HOST http://localhost:8080
ENV SEREVR_RECIPE_ENDPOINT https://s3-eu-west-1.amazonaws.com/test-golang-recipes/%d

EXPOSE 8080
ENTRYPOINT ["/recipe_proxy_server"]
