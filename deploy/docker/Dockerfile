# Stage 1: Build stage
FROM golang:tip-alpine3.22 AS builder

# Install make and other build tools.
RUN apk add --no-cache make

RUN adduser -D -g '' appuser

WORKDIR /build

COPY go.mod go.sum Makefile ./

# Use Makefile to download dependencies.
RUN make deps

COPY . .

# Install Goa framework and tools.
RUN make install-goa

# Build the application using Makefile.
RUN make docker-build

# Stage 2: Final stage
FROM alpine:latest AS deployment

# Create a non-root user.
RUN adduser -D -g '' appuser

COPY --from=builder /build/iam-service /app/main

# Change ownership of the app directory to appuser.
RUN chown -R appuser:appuser /app

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/main"]