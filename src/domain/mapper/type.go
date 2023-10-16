package mapper

import (
	db "src/db/model"
	"src/domain/model"
	api "src/spotifyapi/model"
)

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
