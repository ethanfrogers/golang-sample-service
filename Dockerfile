FROM golang:1.9-alpine as builder

ENV PKG=/go/src/github.com/ethanfrogers/golang-sample-service

ADD . $PKG
WORKDIR $PKG

ENV CGO_ENABLED 0

RUN go install -v

FROM alpine:3.6

RUN apk add -U ca-certificates

COPY --from=builder /go/bin/golang-sample-service /bin/golang-sample-service
ADD ./migrations /bin/migrations

ENTRYPOINT ["/bin/golang-sample-service"]
CMD ["serve"]