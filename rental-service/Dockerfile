FROM golang:1.21.1

WORKDIR /usr/local/go/src/user

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]