FROM golang:1.13 as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

<<<<<<< HEAD
CMD ["/app/main"]
=======
CMD ["/app/cmd/main"]
>>>>>>> 59e458412bf4095ff0ad2e2a9d9433f03ca401f8
