FROM golang:1.20-alpine

WORKDIR /home

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server main.go

EXPOSE 8080
ENTRYPOINT [ "./server" ]