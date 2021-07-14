FROM golang:alpine as builder
RUN go build -o app

FROM golang:alpine
CMD ["./app"]