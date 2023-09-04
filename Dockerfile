FROM golang:1.20.2-alpine
WORKDIR /app
COPY . .
ARG GOARCH=arm64
ARG GOOS=darwin
RUN go build -o fsp ./cmd