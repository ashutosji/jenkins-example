FROM golang:1.21.12-bookworm
WORKDIR /app
COPY . .
RUN go build -v -o /app/main .
CMD ["/app/main"]