FROM golang:alpine 
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./src/*.go ./
RUN  go build -o /main

EXPOSE 8080

CMD [ "/main" ]

