# Prepare
FROM golang:1.18.0-alpine AS prepare
WORKDIR /source
COPY go.mod go.sum /source/
RUN go mod download

# Build
FROM prepare AS build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/deck-api cmd/main.go

# Run
FROM alpine as run
COPY --from=build /source/bin/deck-api /deck-api
ENTRYPOINT ["/deck-api"]
