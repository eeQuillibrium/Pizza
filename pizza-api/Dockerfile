FROM golang:latest

ENV CONFIG_PATH="configs/config.yaml"

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go get ./...

COPY ./ ./
RUN CGO_ENABLED=0 go build -o bin cmd/main.go


CMD ["./bin/main"]