# Use a Rust image as the base image
FROM rust:1.68 AS builder

# Set the working directory to /usr/src/app
WORKDIR /usr/src/myapp

# Copy the Cargo.toml and Cargo.lock files to the container
COPY Cargo.toml Cargo.lock ./

# Copy the source code to the container 
COPY ./ ./
COPY /src/ ./src/
#COPY /src/static ./src/static/

RUN cargo build --release

# Build the application
RUN cargo install --path .

# Create a new stage for the application
FROM debian:buster-slim

# Install OpenSSL, which is required by Actix-web
RUN apt-get update && \
    apt-get install -y openssl && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary from the previous stage to the new stage
WORKDIR /app
COPY --from=builder /usr/local/cargo/bin/RustGloss /app
COPY /static /app/static
COPY /static/style.css /app/static/

# Start the application
CMD ["/app/RustGloss"]
