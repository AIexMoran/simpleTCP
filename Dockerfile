FROM golang:1.15 as builder
WORKDIR /source
COPY ./src /source/src
RUN GOPATH=/source CGO_ENABLED=0 GOOS=linux go build -o server /source/src/mainLogic/*.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /source/server .
CMD ["./server"]