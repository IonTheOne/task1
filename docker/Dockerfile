FROM golang:1.22.0-alpine3.19 as builder
COPY ./task1 /go/task1
WORKDIR /go/task1
RUN go mod vendor
RUN cd ./cmd && \
go build -o main .

FROM alpine:latest
WORKDIR /go/
EXPOSE 8100
COPY --from=builder /go/task1/cmd/main ./app
CMD ["./app", "-listening-address=0.0.0.0:8100"]

