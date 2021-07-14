FROM golang:alpine as builder
COPY . .
RUN go env -w GO111MODULE=off
RUN go build -o /app

FROM golang:alpine
COPY --from=builder /app .
CMD ["./app"]