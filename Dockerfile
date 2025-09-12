# --- build stage ---
FROM golang:1.24-alpine AS build
WORKDIR /src

# modules
ENV CGO_ENABLED=0 GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /out/fund-forge ./cmd/fund-forge

# --- runtime stage ---
FROM alpine:3.20
RUN adduser -D -u 10001 appuser
WORKDIR /app

COPY --from=build /out/fund-forge /app/fund-forge
EXPOSE 8080
USER appuser
ENTRYPOINT ["/app/fund-forge"]
