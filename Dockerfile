# docker-compose up [--build]
FROM golang:alpine3.16 AS build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/myapp
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /go/bin/myapp main.go

FROM scratch
COPY --from=build /go/bin/myapp /go/bin/myapp
ENTRYPOINT ["/go/bin/myapp"]
