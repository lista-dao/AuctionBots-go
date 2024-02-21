FROM golang:1.22

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/app /app/cmd/bot

FROM alpine:3.9

COPY --from=0 /usr/local/bin/app /usr/local/bin/app
COPY ./config /app/config
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["app"]
