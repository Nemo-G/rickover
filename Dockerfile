FROM golang:1.17-alpine AS builder-base

RUN apk update &&\
	apk add	ca-certificates \
	curl \
	git \
	patch \
	make \
	g++ \
	unzip

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM builder-base

WORKDIR /app

COPY . /app
RUN go mod tidy

RUN make build
CMD ["./start-server.sh"]