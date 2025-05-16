package mapper

import db "src/db/model"

func StringToDBGenre(genre string) db.Genre {
	return db.Genre(genre)
}

func DBUserIDGenreMappingFromUserIDAndLabel(userID db.ID, labels []string) []db.UserIDGenreMapping {
	mappings := make([]db.UserIDGenreMapping, len(labels))
	for i, label := range labels {
		mappings[i] = db.UserIDGenreMapping{
			UserID: userID,
			Genre:  StringToDBGenre(label),
		}
	}

	return mappings
}
