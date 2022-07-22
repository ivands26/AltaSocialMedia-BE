FROM golang:1.18

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go build -o main .

CMD ["/app/main"]