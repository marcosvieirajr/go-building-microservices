# syntax=docker/dockerfile:1

# --- first (build) stage ---

# FROM golang:1.17-buster AS build
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

RUN env GOOS=linux GOARCH=arm64 \
    go build -o /app-builded


# --- final (target) stage ---

# FROM gcr.io/distroless/base-debian10
FROM alpine:3.15
LABEL maintainer="Marcos Vieira Jr"

WORKDIR /

COPY --from=build /app-builded /app

ENV HTTP_ADDRESS=:8080 

EXPOSE 8080

# USER nonroot:nonroot

ENTRYPOINT ["/app"]
