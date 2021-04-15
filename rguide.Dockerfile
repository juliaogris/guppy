FROM golang:1.16.3-alpine3.13 as builder
ARG VERSION

WORKDIR /src
COPY go.mod go.sum ./
COPY pkg/rguide/ ./pkg/rguide/
COPY cmd/rguide/ ./cmd/rguide/
RUN go build -ldflags="-X main.version=${VERSION}" ./cmd/rguide

FROM alpine:3.13
COPY --from=builder /src/rguide /app/
ENTRYPOINT /app/rguide --address 0.0.0.0:9090
EXPOSE 9090
