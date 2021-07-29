FROM golang:latest

ENV PORT

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o firego .


CMD ["./firego"]