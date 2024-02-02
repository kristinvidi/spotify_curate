package mapper

import (
	db "src/db/model"
	"src/domain/model"
	api "src/spotifyapi/model"
)

func StringToDBID(id string) db.ID {
	return db.ID(id)
}

func DBIDsToAPIIDs(ids []db.ID) []api.ID {
	var apiIDs []api.ID
	for _, id := range ids {
		apiIDs = append(apiIDs, api.ID(id))
	}
	return apiIDs
}

func IDToDBID(id model.ID) db.ID {
	return db.ID(id)
}
