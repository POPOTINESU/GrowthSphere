FROM golang:1.23.3 AS base
WORKDIR /app

COPY go.mod go.sum ./
RUN apt-get update && \
    apt-get install -y curl netcat-openbsd && \
    go mod download

FROM base AS dev
COPY . .
RUN go install github.com/air-verse/air@latest && \
    chmod +x /app/build/entrypoint.sh

ENTRYPOINT ["/app/build/entrypoint.sh"]
CMD ["air", "-c", "/app/build/.air.toml"]

FROM base AS test
COPY . .
RUN chmod +x /app/build/entrypoint.sh

ENTRYPOINT ["/app/build/entrypoint.sh"]
CMD ["sh", "-c", "go test ./... -v"]

FROM base AS build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/app . &&\
    chmod +x /app/build/entrypoint.sh

FROM build AS production
COPY --from=build /app/app .
ENTRYPOINT ["/app/build/entrypoint.sh"]
CMD ["./app"]