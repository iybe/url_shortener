FROM golang:1.18 as development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8888
RUN go build -o app .
CMD ["./app"]