#!/bin/bash

# Create certs directory if it doesn't exist
mkdir -p ../server/certs

# Generate private key
openssl genrsa -out ../server/certs/server.key 2048

# Generate self-signed certificate
openssl req -new -x509 -sha256 -key ../server/certs/server.key -out ../server/certs/server.crt -days 3650 \
    -subj "/C=US/ST=State/L=City/O=Organization/CN=localhost"

echo "Generated TLS certificates in server/certs/"
echo "  - server.key: Private key"
echo "  - server.crt: Self-signed certificate"
echo ""
echo "To use TLS:"
echo "1. Set use_tls=true in your config.toml"
echo "2. Ensure cert_file and key_file in config.toml point to these files"

# Set appropriate permissions for the private key
chmod 600 ../server/certs/server.key 
