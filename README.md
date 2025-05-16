# Curate: A Spotify app for DJs to discover new music

Welcome to my Spotify playlist curating app - Curate!

## Features
- Labeling: Categorize your artists by genre.
- Recent in Genre: Configure regularly generating playlists containing your followed artists' new albums.

## Upcoming Features
- Hidden gems: Surface your favourite artists' lesser-known music.
- Related Artists: Select an artist to use as reference and generate a playlist based on their related artists.
- Compilations: See compilation albums that your artists are included on (great for discovering similar artists!).

## Setup Instructions

### Prerequisites (Mac Only)
- Install [Homebrew](https://docs.brew.sh/Installation).
- Install [Docker Desktop](https://www.docker.com/get-started/).
- Install [Postman](https://www.postman.com/downloads/).

### Service Setup
1. **Database**: See the README in the `postgres` directory for database setup instructions.
2. **Backend**: See the README in the `src` directory for backend application setup.
3. 

### Running the App
- Ensure Docker is running and the `spotify_db` container is running. If not, check the Database setup above.
- Run the gRPC server:
    - `cd src`
    - `go run main.go`
- Start up Postman.
