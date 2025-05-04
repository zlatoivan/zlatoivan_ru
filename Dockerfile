FROM golang:1.23.5-alpine

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o zlatoivan cmd/server/main.go

CMD ["./zlatoivan"]