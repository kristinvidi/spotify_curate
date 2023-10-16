package query

import (
	"context"
	"src/db/model"
	"time"
)

func (p *PostgresDB) GetLastCreatedAtDateForPlaylistOfGenre(userID model.ID, genre model.Genre) (*time.Time, error) {
	var lastCreated time.Time
	err := p.db.NewSelect().
		Model(&lastCreated).
		TableExpr("playlist_recent_in_genre_generated_status priggs").
		ColumnExpr("max(created_at)").
		Join("inner join user_spotify_id_genre_mapping usigm on usigm.id = priggs.genre_id").
		Where("usigm.user_spotify_id = ?", userID).
		Where("usigm.genre = ?", genre).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return &lastCreated, nil
}

func (p *PostgresDB) InsertPlaylistRecentInGenreGeneratedStatus(status model.PlaylistRecentInGenreGeneratedStatus) error {
	return p.insertNoConflict(&status)
}
