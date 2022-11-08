FROM golang:1.17

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
        -o /usr/local/bin/app \
        /app

FROM alpine:3.9

COPY --from=0 /usr/local/bin/app /usr/local/bin/app
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["app"]
