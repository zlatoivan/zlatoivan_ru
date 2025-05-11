FROM golang:1.23.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o zlatoivan cmd/server/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app /app
CMD ["/app/zlatoivan"]