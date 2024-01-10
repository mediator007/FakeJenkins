FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=1

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]