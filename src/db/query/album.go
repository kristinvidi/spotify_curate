package query

import (
	"context"
	"time"

	"src/db/model"
)

func (p *PostgresDB) GetAlbumIDsForGenreAfterDate(userID model.ID, genre model.Genre, afterDate time.Time) ([]model.ID, error) {
	var albumIDs []model.ID
	err := p.db.NewSelect().
		Model(&albumIDs).
		TableExpr("spotify_user u").
		Column("aasim.album_spotify_id").
		Join("inner join user_spotify_id_genre_mapping usigm on u.spotify_id = usigm.user_spotify_id").
		Join("inner join user_artist_spotify_id_genre_mapping uasigm on u.spotify_id = uasigm.user_spotify_id and usigm.id = uasigm.genre_id").
		Join("inner join artist_album_spotify_id_mapping aasim on uasigm.artist_spotify_id = aasim.artist_spotify_id").
		Join("inner join spotify_album a on aasim.album_spotify_id = a.spotify_id").
		Where("u.spotify_id = ?", userID).
		Where("usigm.genre = ?", genre).
		Where("a.release_date > ?", model.FormatPostgresTime(afterDate, model.TimeFormatPostgresDate)).
		Where("a.album_type != 'compilation'").
		Scan(context.Background())

	return albumIDs, err
}
