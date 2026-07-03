# ---------- build stage ----------
FROM golang:1.24-alpine AS builder
WORKDIR /app

# deps
COPY go.mod go.sum ./
RUN go mod download

# source
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go


# ---------- runtime stage ----------
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates curl

# نصب migrate CLI
RUN curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz \
    | tar xvz && mv migrate /usr/local/bin/migrate

# کپی باینری اپ
COPY --from=builder /app/app .

# کپی فولدر migration ها
COPY --from=builder /app/database/migrations ./database/migrations

EXPOSE 3000

#CMD sh -c 'migrate -path database/migrations -database "$DATABASE_URL" up || true && ./app'

CMD sh -c 'migrate -path database/migrations -database "$DATABASE_URL" up && ./app'