#Build stage

FROM golang:1.17-alpine3.16 AS builder
WORKDIR /
COPY . .
RUN go build -o sample main.go

# Run stage

FROM alpine:3.16
WORKDIR /
COPY --from=builder /sample .

EXPOSE 10000
CMD ["./sample"]




