# Go - Compile and execute go http web server
FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 3000
ENTRYPOINT ["/app/main"]
