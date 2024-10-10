# stage 1
FROM golang:1.23-alpine as build

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -mod vendor -tags musl -o  server ./cmd/server/server.go


# stage 2
FROM alpine:3.12
# Install bash
RUN apk add --no-cache bash

WORKDIR /app

COPY bin/run bin/run

COPY --from=build /build/server .

RUN  chmod +x bin/run
EXPOSE 8081

CMD ["bin/run", "server"]
