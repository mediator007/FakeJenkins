FROM golang:alpine

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go build -o main main.go

CMD ["./main"]