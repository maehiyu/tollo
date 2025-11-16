FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download

COPY . .

RUN go build -o /usr/local/bin/chatservice ./cmd/chatservice

EXPOSE 50052

ENTRYPOINT ["chatservice"]
