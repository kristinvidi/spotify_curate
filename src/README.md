# Backend Setup

## Development Setup
1. Install `go` using [Homebrew go installation](https://formulae.brew.sh/formula/go)
    - *See `go.mod` for current `go` version*
    - Verify installation by running `go version`
2. Install dependencies: `go mod download`
3. Copy the example configuration file to create your own: `cp config/config.toml.example config/config.toml`.
4. Edit `config.toml` and add your Spotify API credentials:
   - `client_id`: Your Spotify application client ID
   - `client_secret`: Your Spotify application client secret

## Running the Application
1. Make sure you're in the `src` directory
2. Run the application: `go run main.go`

### Setting up Postman for gRPC
1. Open Postman.
2. In the top bar select `New` -> `gRPC`.
3. Under `URL`, put `localhost:50051`.
4. Import the `proto` definition:
    1. Under `Select a method`, select `Import a .proto file`.
    2. Select the file from `proto/spotify_curate.proto`.
    3. Select `Import as API`.
    4. Name the API and click `Create a New API`.
    5. Now when you click `Select a method` you should see a list of the available endpoints.
5. To call an endpoint:
    1. Select the desired endpoint from `Select a method`.
    2. Ensure you are on the `Message` tab.
    3. Select `Use Example Message` to see the necessary message structure.
