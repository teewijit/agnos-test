FROM golang:1.23-alpine

WORKDIR /api

COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]