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

Naming shortcuts:
- If names include several spotify IDs, these are simplified by chainining the entity names followed by `spotify_id`. For example, `user_spotify_id_artist_spotify_id` would be simplified to `user_artist_spotify_id`.