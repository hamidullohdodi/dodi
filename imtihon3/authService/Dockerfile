FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o authservice ./cmd/main.go

EXPOSE 8070

CMD ["/app/authservice"]
