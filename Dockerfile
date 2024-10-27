FROM golang:1.22

LABEL authors="ivan"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build cmd/server/main.go -o zlatoivan_ru

CMD ["./zlatoivan_ru"]
