FROM golang:1.24-alpine3.22
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/query
WORKDIR /go/src/query
COPY . /go/src/query
RUN go mod download
EXPOSE 8083
CMD ["go", "run", "cmd/server/main.go"]
