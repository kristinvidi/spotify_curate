# Postgres Database

## Setting up and running the database

### Initial Setup
1. Make sure Docker Desktop is running
    - From the `postgres` directory, run: `docker-compose build`.

### Running the database an applying schema changes
You have several options to run the database and apply schema changes:

1. Run both database and apply schema changes:
    - Run `docker-compose up`.
    - This will start both the Postgres database and run Liquibase migrations.

2. Run only the database (in detached mode):
   - Run `docker-compose up -d spotify_db`.
   - Use this when you just need the database running without applying new schema changes.

3. Apply schema changes using the script:
   - Run `./apply_schema.sh`.
   - Use this to apply new database migrations after making schema changes.

4. Stop all containers:
   - Run `docker-compose down`.
   - Use this when you're done and want to stop the database.

## Guidelines for naming schema objects
Examples below are based on the `user_artist_spotify_id_mapping` table.

 Object | Format | Example
 :--- | :--- | :---
 Primary keys | `pk_table_name_id` | `pk_user_artist_spotify_id_mapping_id`
 Foreign keys | `fk_table_name_column_name` | `fk_user_artist_spotify_id_mapping_user_spotify_id`
 Unique keys | `unique_table_name_column_names` | `unique_user_artist_spotify_id_mapping_user_artist_spotify_id`

Notes on naming:
- The `spotify_user` table was named as such to avoid conflicts with the built-in `user` table. All references to this table in mapping tables etc (like `user_spotify_id_artist_spotify_id`) simplify this to `user`.
- Given the `spotify_user` edge-case, all Spotify entities are prefaced by `spotify` (like `spotify_artist` and `spotify_album`). This also ensures all these tables are grouped when previewing the database contents via the drop-down.
- If names include several spotify IDs, these are simplified by chainining the entity names followed by `spotify_id`. For example, `user_spotify_id_artist_spotify_id` would be simplified to `user_artist_spotify_id`.
