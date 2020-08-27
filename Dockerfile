FROM golang:1.15 as builder
WORKDIR /source
COPY ./ /source/
RUN CGO_ENABLED=0 GOOS=linux go build -o server /source/cmd/server/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /source/server .
CMD ["./server"]