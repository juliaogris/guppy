FROM golang:1.16.3-alpine3.13 as builder
ARG VERSION

WORKDIR /src
COPY go.mod go.sum ./
COPY pkg/echo/ ./pkg/echo/
COPY cmd/echo/ ./cmd/echo/
RUN go build -ldflags="-X main.version=${VERSION}" ./cmd/echo

FROM alpine:3.13
COPY --from=builder /src/echo /app/
ENTRYPOINT /app/echo --address 0.0.0.0:9090
EXPOSE 9090
