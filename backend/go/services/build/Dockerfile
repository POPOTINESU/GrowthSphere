FROM golang:1.23.3 AS base
WORKDIR /app

COPY go.mod go.sum ./
RUN apt-get update && \
    apt-get install -y curl netcat-openbsd postgresql-client && \
    go mod download && \
    rm -rf /var/lib/apt/lists/*

COPY build/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh
ENTRYPOINT ["/app/build/entrypoint.sh"]

FROM base AS development
COPY . .
RUN go install github.com/air-verse/air@latest && \
    curl -sSf https://atlasgo.sh | sh && \
    rm -rf /var/lib/apt/lists/*
CMD ["air", "-c", "/app/build/.air.toml"]

FROM base AS build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/app .

FROM base AS production
COPY --from=build /app/app .
CMD ["./app"]