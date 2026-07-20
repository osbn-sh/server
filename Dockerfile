# ---------- build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

# ---------- runtime stage ----------
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/database/migrations ./database/migrations

EXPOSE 3000

CMD ["./app"]