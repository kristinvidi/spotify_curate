# Curate: A Spotify App for DJ Music Discovery

Welcome to my Spotify playlist curating app - Curate!

## Features
- Labeling: Categorize your artists by genre.
- Recent in Genre: Configure regularly generating playlists containing your followed artists' new albums.

## Upcoming Features
- Hidden gems: Surface your favourite artists' lesser-known music.
- Related Artists: Select an artist to use as reference and generate a playlist based on their related artists.
- Compilations: See compilation albums that your artists are included on (great for discovering similar artists!).

## Repo Setup

### Install things!
- Install Docker
- Install Go. See `src/go.mod` for current Go version.
- Install repo packages. From the `src` directory, run `go mod download`.

### Set up local database
Head to the README in the `postgres` directory for more info!

### Running the app
- At time of writing, there are no endpoints. You can configure the job you want by specifying the `job` variable in `src/main.go`.
- First you should:
    - Run `job` as `UPDATE_USER_DATA` to fetch your user profile, followed artists, and their albums. This takes a while! Be patient.
    - Update the `user_spotify_id_genre_mapping` table with your chosen genres. See the sample SQL at `/postgres/sample_sql/mock_insert_user_spotify_id_genre_mapping.sql` for reference.
        - Pro tip - you can dump a csv of the `spotify_artist` table into a spreadsheet and label them that way!
    - Update the `user_spotify_id_genre_mapping` table with your artist to genre mappings. See the sample SQL at `/postgres/sample_sql/mock_insert_user_artist_spotify_id_genre_mapping.sql` for reference.
    - Run `job` as `CREATE_PLAYLIST_RECENT_IN_GENRE`, with the genre you want specified by the `genre` variable.