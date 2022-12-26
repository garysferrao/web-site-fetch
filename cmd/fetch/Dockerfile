# docker build -f cmd/fetch/Dockerfile .
FROM golang:1.19 AS build

WORKDIR /tmp/web-site-fetch
COPY . /tmp/web-site-fetch
RUN go mod download
RUN go build -o fetch cmd/fetch/fetch.go

FROM golang:1.19
COPY --from=build /tmp/web-site-fetch/fetch fetch
