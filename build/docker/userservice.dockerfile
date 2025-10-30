FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download

COPY . .

RUN go build -o /usr/local/bin/userservice ./cmd/userservice

EXPOSE 50051

ENTRYPOINT ["userservice"]
