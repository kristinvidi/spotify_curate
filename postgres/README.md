# Postgres Database

## Starting the docker container
- From the `postgres` directory, run `docker-compose build`.

## Applying the schema
- From the `postgres` directory, run `.\apply_schema.sh`.

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