FROM golang:1.22

LABEL authors="ivan"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o zlatoivan cmd/server/main.go

CMD ["./zlatoivan"]

EXPOSE 7072