FROM golang:1.19 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -o outboundSentWriter .

FROM scratch
COPY --from=builder /build/outboundSentWriter /
COPY --from=builder /build/outboundSentWriter.log /

ENTRYPOINT ["/outboundSentWriter"]