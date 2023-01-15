FROM golang:latest

ENV GO111MODULE=on

RUN apt-get update && apt-get install -y postgresql-client

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chmod +x seed-db.sh

CMD ["./seed-db.sh"]

RUN go build -o main .

EXPOSE 3000

ENTRYPOINT ["/app/main"]