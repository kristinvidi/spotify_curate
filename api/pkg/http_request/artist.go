package httprequest

import (
	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/converter"
	"spotify_app/api/pkg/model"
)

func (h *HttpRequest) GetArtistTopTracks(accessToken apptype.AccessToken, artistID string) (*model.Tracks, error) {
	countryCode, err := h.configManager.GetUserCountryCode()
	if err != nil {
		return nil, err
	}

	request, err := converter.BuildGetArtistTopTracksRequest(accessToken, artistID, *countryCode)
	if err != nil {
		return nil, err
	}

	response, err := h.DoRequestAndCheckStatusIsOK(request)
	if err != nil {
		return nil, err
	}

	convertedResponse, err := converter.DecodeGetArtistTopTracksResponse(*response)
	if err != nil {
		return nil, err
	}

	return &convertedResponse.Tracks, nil
}

func (h *HttpRequest) GetArtistRelatedArtists(accessToken apptype.AccessToken, artistID string) (*model.Artists, error) {
	request, err := converter.BuildGetArtistRelatedArtistsRequest(accessToken, artistID)
	if err != nil {
		return nil, err
	}

	response, err := h.DoRequestAndCheckStatusIsOK(request)
	if err != nil {
		return nil, err
	}

	convertedResponse, err := converter.DecodeGetArtistRelatedArtistsResponse(*response)
	if err != nil {
		return nil, err
	}

	return &convertedResponse.Artists, nil
}

// func (h *HttpRequest) GetArtistAlbums(accessToken apptype.AccessToken, artistID string) (*model.Albums, error) {
// 	countryCode, err := h.configManager.GetUserCountryCode()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var albums []model.Album
// 	var after *string
// 	total := 1

// 	for len(albums) < total {
// 		request, err := converter.BuildGetArtistsAlbumsRequest(accessToken, artistID, after, *countryCode)
// 		if err != nil {
// 			return nil, err
// 		}

// 		response, err := h.DoRequestAndCheckStatusIsOK(request)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer response.Body.Close()

// 		albumResponse, err := converter.DecodeGetArtistsAlbumsResponse(*response)
// 		if err != nil {
// 			return nil, err
// 		}

// 		albums = append(albums, )

// 		// artistItems := artistResponse.Artists
// 		// followedArtists = append(followedArtists, artistItems.ArtistList...)

// 		// after = &artistItems.Cursors.After
// 		// total = artistItems.Total
// 	}

// 	return followedArtists, nil
// }
