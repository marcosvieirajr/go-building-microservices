# syntax=docker/dockerfile:1

# --- first (build) stage ---

FROM golang:1.17-buster AS build
# FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -v -o /app

# --- final (target) stage ---

FROM gcr.io/distroless/base-debian10
# FROM alpine:3.15

WORKDIR /

# RUN apk update && apk add wget && \
#     wget https://github.com/fullstorydev/grpcurl/releases/download/v1.8.6/grpcurl_1.8.6_linux_x86_64.tar.gz && \
#     gzip -d grpcurl_1.8.6_linux_x86_64.tar.gz && \
#     tar -xf grpcurl_1.8.6_linux_x86_64.tar && \
#     rm grpcurl_1.8.6_linux_x86_64.tar 

COPY --from=build /app /app

# ENV HTTP_ADDRESS=:8080

EXPOSE 9092

USER nonroot:nonroot

ENTRYPOINT ["/app"]
