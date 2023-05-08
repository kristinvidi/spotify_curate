package httprequest

import (
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/converter"
	"spotify_app/api/pkg/model"
)

func (h *HttpRequest) GetFollowedArtists(accessToken apptype.AccessToken) ([]model.Artist, error) {
	var followedArtists []model.Artist
	var after *string
	total := 1

	for len(followedArtists) < total {
		request, err := converter.BuildGetFollowedArtistsRequest(accessToken, after)
		if err != nil {
			return nil, err
		}

		response, err := h.DoRequestAndCheckStatus(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		artistResponse, err := converter.DecodeGetFollowedArtistsResponse(*response)
		if err != nil {
			return nil, err
		}

		artistItems := artistResponse.Artists
		followedArtists = append(followedArtists, artistItems.ArtistList...)

		after = &artistItems.Cursors.After
		total = artistItems.Total
	}

	return followedArtists, nil
}
