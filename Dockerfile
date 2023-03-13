FROM golang:1.20-alpine

ENV GOPATH=/

COPY . .

RUN go mod download
RUN go build -o go-todo-app ./cmd/main.go

CMD ["./go-todo-app"]
