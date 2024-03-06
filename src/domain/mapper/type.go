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

func APIIDToDBID(id api.ID) db.ID {
	return db.ID(id)
}

func IDToDBID(id model.ID) db.ID {
	return db.ID(id)
}

func URIsToAPIURIs(uris []model.URI) []api.URI {
	var apiURIs []api.URI
	for _, uri := range uris {
		apiURIs = append(apiURIs, api.URI(uri))
	}
	return apiURIs
}

func IDsToDBIDs(ids []model.ID) []db.ID {
	var dbIDs []db.ID
	for _, id := range ids {
		dbIDs = append(dbIDs, IDToDBID(id))
	}
	return dbIDs
}
