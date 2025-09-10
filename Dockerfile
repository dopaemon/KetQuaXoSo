# ===== Stage 1: Build binary =====
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags headless -o KetQuaXoSo .

# ===== Stage 2: Runtime =====
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/KetQuaXoSo /app/KetQuaXoSo

EXPOSE 8080

CMD ["/app/KetQuaXoSo", "--api"]
