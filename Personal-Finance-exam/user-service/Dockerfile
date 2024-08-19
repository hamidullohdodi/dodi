FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service ./cmd/main.go

EXPOSE 8070

CMD ["/app/user-service"]

# FROM golang:1.22.5-alpine

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY . .

# RUN go build -o api-gateway ./cmd/main.go 

# EXPOSE 8080

# CMD ["/app/api-gateway"]  