# gRPC Server

This directory contains the gRPC server implementation for the Spotify Curate application.

## Configuration

The server can be run in either secure (TLS) or insecure mode, configured via `config.toml`.

### Insecure Mode (Development)
For local development, you can run the server in insecure mode by setting the following in your config:
- Set `use_tls = false` in the `[grpc]` section of your config.

### Secure Mode (TLS)
For production or when security is required:

1. Generate TLS certificates by running: `./scripts/generate_certs.sh` from the `server` directory.
   This will create:
   - `certs/server.key`: Private key
   - `certs/server.crt`: Self-signed certificate

2. Update your config:
   - Set `use_tls = true` in the `[grpc]` section
   - Ensure `cert_file = "server/certs/server.crt"` in the `[grpc.tls]` section
   - Ensure `key_file = "server/certs/server.key"` in the `[grpc.tls]` section

### Security Notes
- Never commit certificates to version control
- Keep private keys secure (`server.key` has 600 permissions)
- Use proper CA-signed certificates for production
- Self-signed certificates are only suitable for development

## Client Configuration

### Postman
When TLS is enabled:
1. Import the server certificate (`server.crt`)
2. Enable TLS in the request settings
3. Set "Server Certificate Validation" according to your needs:
   - OFF for self-signed certificates
   - ON for CA-signed certificates 
