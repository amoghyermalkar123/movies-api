FROM golang:latest
RUN mkdir /app
ADD go.mod /app/
ADD go.sum /app/
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
