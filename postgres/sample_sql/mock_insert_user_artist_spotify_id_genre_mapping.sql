insert into user_artist_spotify_id_genre_mapping (user_spotify_id, artist_spotify_id, genre_id) values 
('user_spotify_id', 'artist_spotify_id', (select id from user_spotify_id_genre_mapping where genre = 'Genre1' and user_spotify_id = 'user_spotify_id')),
('user_spotify_id', 'artist_spotify_id', (select id from user_spotify_id_genre_mapping where genre = 'Genre2' and user_spotify_id = 'user_spotify_id'));