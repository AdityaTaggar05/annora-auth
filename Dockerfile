# ---------- Build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# IMPORTANT: correct build path
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -buildid=" -trimpath \
    -o auth-service ./cmd/auth-service

# ---------- Runtime stage ----------
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/auth-service .
COPY ./keys ./keys
COPY ./internal/mailer/templates ./internal/mailer/templates

EXPOSE 8080

CMD ["./auth-service"]
