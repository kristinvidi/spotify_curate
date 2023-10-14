package mapper

import db "src/db/model"

func StringToDBGenre(genre string) db.Genre {
	return db.Genre(genre)
}
