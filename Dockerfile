# ───────────────────────────────
# Stage 1: Base (common setup)
# ───────────────────────────────
FROM golang:1.25 AS base
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ───────────────────────────────
# Stage 2: Development (with Air)
# ───────────────────────────────
FROM base AS dev
RUN go install github.com/air-verse/air@latest
ENV PATH="/go/bin:${PATH}"
CMD ["air"]

# ───────────────────────────────
# Stage 3: Builder (compile binary)
# ───────────────────────────────
FROM base AS builder
RUN go build -o social ./cmd/api

# ───────────────────────────────
# Stage 4: Production (slim runtime)
# ───────────────────────────────
FROM debian:bookworm-slim AS prod
WORKDIR /app
COPY --from=builder /app/social .
EXPOSE 8080
CMD ["./social"]
