package mapper

import (
	db "src/db/model"
	api "src/spotifyapi/model"
)

func DBIDsToAPIIDs(ids []db.ID) []api.ID {
	var apiIDs []api.ID
	for _, id := range ids {
		apiIDs = append(apiIDs, api.ID(id))
	}
	return apiIDs
}
