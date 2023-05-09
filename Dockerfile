FROM golang:1.20-alpine3.17

RUN apk add --no-cache make

WORKDIR /myapp

COPY . .

RUN make build

EXPOSE 8080

CMD ["./build/fizzbuzz"]