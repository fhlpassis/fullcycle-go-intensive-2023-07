FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o api cmd/api/main.go

CMD ["./api"]


FROM scratch

COPY --from=builder /app/api /

CMD ["/api"]