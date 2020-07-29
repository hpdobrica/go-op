FROM golang:alpine AS build

ENV OP_CLI_VERSION=v1.3.0

RUN apk --no-cache add curl
RUN curl -sS -o 1password.zip https://cache.agilebits.com/dist/1P/op/pkg/$OP_CLI_VERSION/op_linux_386_$OP_CLI_VERSION.zip
RUN head -n 5 1password.zip
RUN unzip 1password.zip op -d /usr/bin
RUN rm 1password.zip


RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get -d -v
RUN go build -o /go/bin/app

FROM scratch
# FROM alpine:3.12

COPY --from=build /usr/bin/op /usr/bin/
COPY --from=build /go/bin/app /go/bin/app

CMD ["/go/bin/app"]
