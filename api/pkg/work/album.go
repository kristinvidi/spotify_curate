package work

// func StoreAlbumInfo(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken) (model.Artists, error) {
// 	db := datastore.NewArtistTextDB()
// 	artists, err := db.ReadArtistIDToArtist()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var albumIDs []string
// 	for a := range artists {
// 		artistAlbums, err := httpRequest.GetArtistAlbums(accessToken, a)
// 		if err != nil {
// 			return nil, err
// 		}

// 		albumIDs = append(albumIDs, artistAlbums.IDs()...)
// 	}

// 	// Store albums

// 	return nil, nil
// }
