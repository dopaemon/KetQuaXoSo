# ===== Stage 1: Build binary =====
FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y --no-install-recommends \
    libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev \
    libgl1-mesa-dev xorg-dev \
    && rm -rf /var/lib/apt/lists/*

RUN go mod tidy
RUN go build -v

# ===== Stage 2: Runtime =====
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/KetQuaXoSo /app/KetQuaXoSo

RUN apt-get update && apt-get install -y --no-install-recommends \
    libgl1-mesa-glx libx11-6 libxrandr2 libxcursor1 libxinerama1 libxi6 \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 8080

CMD ["/app/KetQuaXoSo", "--api"]
