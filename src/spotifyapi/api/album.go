package api

import (
	"src/spotifyapi/convert"
	"src/spotifyapi/model"
)

type Album struct {
	api            *API
	albumConverter *convert.Album
}

func NewAlbum(api *API, albumConverter *convert.Album) *Album {
	return &Album{
		api:            api,
		albumConverter: albumConverter,
	}
}

func (a *Album) GetAlbumTracks(albumIDs []model.ID) ([]*model.GetAlbumTracksResponse, error) {
	var responses []*model.GetAlbumTracksResponse

	for _, albumID := range albumIDs {
		offset := 0
		trackLen := 0
		total := 1
		batchSize := 50
		for trackLen < total {
			inputs := model.NewRequestInput(&albumID, nil, nil, &offset, &batchSize, nil)
			response, err := a.api.DoRequest(a.albumConverter.BuildGetAlbumTracksRequest, inputs)
			if err != nil {
				return nil, err
			}

			defer response.Body.Close()

			decodedResponse, err := a.albumConverter.DecodeGetAlbumTracksResponse(*response)
			if err != nil {
				return nil, err
			}

			// set total number of tracks
			total = decodedResponse.Total

			// set let of processed tracks
			trackLen += len(decodedResponse.Tracks)

			// add response to responses
			responses = append(responses, decodedResponse)

			offset = trackLen
		}
	}

	return responses, nil
}
