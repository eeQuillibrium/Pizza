FROM golang:latest

ENV CONFIG_PATH="configs/config.yaml"
ENV DB_PASSWORD="secret"

COPY ./ ./
RUN go mod download

RUN go build -o ./bin ./cmd/main.go

CMD ["./bin/main"]
