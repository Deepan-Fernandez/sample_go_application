FROM golang:1.17-alpine3.16
WORKDIR /
COPY . .
RUN go build -o sample main.go

EXPOSE 10000

CMD ["./sample"]
