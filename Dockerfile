FROM golang:1.17.3 as gobuild
ARG VERSION=latest

WORKDIR /Users/twitlin/Code/repos/coolguy1771/rest-api/
ADD go.mod go.sum main.go ./
ADD pkg ./pkg
ADD docs ./docs

RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -mod=vendor -o go-api -ldflags "-X main.version=$VERSION" main.go

FROM gcr.io/distroless/base

COPY --from=gobuild /Users/twitlin/Code/repos/coolguy1771/rest-api/ /bin

ENTRYPOINT ["/bin/go-api"]