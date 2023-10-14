package api

import (
	"src/spotifyapi/convert"
	"src/spotifyapi/model"
)

type Playlist struct {
	api               *API
	playlistConverter *convert.Playlist
}

func NewPlaylist(api *API, playlistConverter *convert.Playlist) *Playlist {
	return &Playlist{
		api:               api,
		playlistConverter: playlistConverter,
	}
}

func (p *Playlist) CreatePlaylist(userID model.ID, name string, public, collaborative bool, description string) (*model.CreatePlaylistResponse, error) {
	playlistInputs := model.NewCreatePlaylistInputs(name, public, collaborative, description)
	requestInputs := model.NewRequestInput(&userID, nil, nil, nil, nil, playlistInputs)

	response, err := p.api.DoRequest(p.playlistConverter.BuildCreatePlaylistRequest, requestInputs)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	decodedResponse, err := p.playlistConverter.DecodeCreatePlaylistResponse(*response)
	if err != nil {
		return nil, err
	}

	return decodedResponse, nil
}

func (p *Playlist) AddTracksToPlaylist(playlistID model.ID, trackIDs []model.URI) ([]*model.AddTracksToPlaylistResponse, error) {
	var responses []*model.AddTracksToPlaylistResponse

	batchSize := 100
	for len(trackIDs) > 0 {
		if len(trackIDs) < batchSize {
			batchSize = len(trackIDs)
		}

		trackSubset := trackIDs[:batchSize]
		trackIDs = trackIDs[batchSize:]
		requestInputs := model.NewRequestInput(&playlistID, trackSubset, nil, nil, &batchSize, nil)

		response, err := p.api.DoRequest(p.playlistConverter.BuildAddTracksToPlaylistRequest, requestInputs)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		decodedResponse, err := p.playlistConverter.DecodeAddTracksToPlaylistResponse(*response)
		if err != nil {
			return nil, err
		}

		responses = append(responses, decodedResponse)
	}

	return responses, nil
}
