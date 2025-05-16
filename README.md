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
- Install [DBeaver](https://dbeaver.io/download/).

### Service Setup
1. **Database**: See the README in the `postgres` directory for database setup instructions.
2. **Backend**: See the README in the `src` directory for backend application setup.

### Running the App
- Ensure Docker is running and the `spotify_db` container is running. If not, check the `Database` setup above.
- Run the gRPC server per below. See `Backend` instructions below if required for Postman setup.
    - `cd src`
    - `go run main.go`
- Interact with Postman! To get started you will want to:
    - Run the `AuthenticateUser` endpoint to authenticate with Spotify and get your UserID.
    - Run the `UpdateUserData` endpoint.
    - Take note of your Spotify User ID when it returns.
    - Create your labels using the `CreateLabelsForUser` endpoint. You will need your User ID here.
    - Run the `GetUnmappedArtistsForUser` endpoint to view which artists need to be associated to labels.
    - Associate your followed artists to labels using the `CreateGenreToArtistMappings` endpoint.
    - *Remember that you can click on `Use Example Message` in the Postman console to view the structure of the API request.*
