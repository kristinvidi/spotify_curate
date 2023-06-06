package httprequest

import (
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/converter"
	"spotify_app/api/pkg/model"
)

func (h *HttpRequest) CreatePlaylist(accessToken apptype.AccessToken, playlistName string) (*apptype.PlaylistID, error) {
	userID, err := h.configManager.GetUserID()
	if err != nil {
		return nil, err
	}

	input := model.NewCreatePlaylistRequest(playlistName, false, false, "")

	request, err := converter.BuildCreatePlaylistRequest(accessToken, *userID, *input)
	if err != nil {
		return nil, err
	}

	response, err := h.DoRequestAndCheckStatusIsOK(request)
	if err != nil {
		return nil, err
	}

	res, err := converter.DecodeCreatePlaylistResponse(*response)
	if err != nil {
		return nil, err
	}

	return &res.ID, nil
}

func (h *HttpRequest) AddTracksToPlaylist(accessToken apptype.AccessToken, playlistID apptype.PlaylistID, tracks model.Tracks) error {
	requests, err := converter.BuildAddTracksToPlaylistRequests(accessToken, playlistID, tracks)
	if err != nil {
		return err
	}

	for _, r := range *requests {
		_, err = h.DoRequestAndCheckStatusIsOK(&r)
		if err != nil {
			return err
		}
	}

	return nil
}
