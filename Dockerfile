# ===== Stage 1: Build binary =====
FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y --no-install-recommends \
    libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev \
    libgl1-mesa-dev xorg-dev \
    && rm -rf /var/lib/apt/lists/*

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags headless -v .

# ===== Stage 2: Runtime =====
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/KetQuaXoSo /app/KetQuaXoSo

EXPOSE 8080

CMD ["/app/KetQuaXoSo", "--api"]
