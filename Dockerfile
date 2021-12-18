FROM golang:alpine
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .
RUN go mod download && go build -o main ./src

EXPOSE 80

CMD ["./main"]