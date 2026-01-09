# Curate: A Spotify app for DJs to discover new music

Welcome to my Spotify playlist curating app - Curate!

## Features
- **Frontend UI**: A premium dark-mode web application to manage your library.
- **Labeling**: Categorize your artists by "Artist Tags" (Genres).
- **Recent in Genre**: Configure regularly generating playlists containing your followed artists' new albums.

## Upcoming Features
- Hidden gems: Surface your favourite artists' lesser-known music.
- Related Artists: Select an artist to use as reference and generate a playlist based on their related artists.
- Compilations: See compilation albums that your artists are included on (great for discovering similar artists!).

## Setup Instructions

### Prerequisites
- Install [Homebrew](https://docs.brew.sh/Installation).
- Install [Docker Desktop](https://www.docker.com/get-started/).
- Install [Go](https://go.dev/doc/install) (for backend).
- Install [Node.js](https://nodejs.org/en/download/) (for frontend).
- Install [DBeaver](https://dbeaver.io/download/) (optional, for DB inspection).

### Service Setup
1. **Database**: See the README in the `postgres` directory for database setup instructions.
2. **Backend**: See the README in the `src` directory for backend application setup.
3. **Frontend**: The frontend is located in the `web` directory. Run `npm install` in `web` to install dependencies.

### Running the App

We use a `Makefile` to simplify running the application.

- **Start All Services**:
    ```bash
    make run
    ```
    This starts both the backend (gRPC server) and frontend (Next.js) concurrently. Open [http://localhost:3000](http://localhost:3000) to use the app.

- **Stop All Services**:
    ```bash
    make stop
    ```

- **Generate Proton Files**:
    ```bash
    make proto
    ```
    Regenerates Go protobuf files and copies the definition to the frontend.

### Manual Usage (Legacy)
If you prefer running services manually or using Postman:
- Ensure Docker is running and the `spotify_db` container is running.
- **Backend**: `cd src && go run main.go`
- **Frontend**: `cd web && npm run dev`
