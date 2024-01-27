package api

import (
	"fmt"

	"src/spotifyapi/convert"
	"src/spotifyapi/model"

	"go.openly.dev/pointy"
)

type User struct {
	api             *API
	userConverter   *convert.User
	artistConverter *convert.Artist
}

func NewUser(api *API, userConverter *convert.User, artistConverter *convert.Artist) *User {
	return &User{api: api, userConverter: userConverter, artistConverter: artistConverter}
}

func (u *User) GetCurrentUsersProfile() (*model.GetCurrentUsersProfileResponse, error) {
	response, err := u.api.DoRequest(u.userConverter.BuildCurrentUsersProfileRequest, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decodedResponse, err := u.userConverter.DecodeCurrentUsersProfileResponse(*response)
	if err != nil {
		return nil, err
	}

	return decodedResponse, nil
}

func (u *User) GetCurrentUsersFollowedArtists() ([]*model.GetFollowedArtistsResponse, error) {
	var responses []*model.GetFollowedArtistsResponse

	var after *string
	artistLen := 0
	total := 1
	batchSize := 50
	for artistLen < total {
		inputs := model.NewRequestInput(nil, nil, after, nil, &batchSize, nil)
		response, err := u.api.DoRequest(u.userConverter.BuildGetFollowedArtistsRequest, inputs)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		decodedResponse, err := u.userConverter.DecodeGetFollowedArtistsResponse(*response)
		if err != nil {
			return nil, err
		}

		// set after to the last artist in the list
		after = pointy.String(decodedResponse.Artists.Cursors.After)

		// set total number of artists
		total = decodedResponse.Artists.Total

		// set let of processed artists
		artistLen += len(decodedResponse.Artists.ArtistList)

		// add response to responses
		responses = append(responses, decodedResponse)
	}

	return responses, nil
}

func (u *User) GetCurrentUsersFollowedArtistsToAlbums(artists model.Artists) (map[model.ID][]*model.GetArtistsAlbumsResponse, error) {
	idToResponses := make(map[model.ID][]*model.GetArtistsAlbumsResponse)

	for i, a := range artists {
		fmt.Printf("%d of %d\n", i, len(artists))
		responsesForArtist, err := u.getArtistsAlbumsForArtist(a)
		if err != nil {
			return nil, err
		}

		idToResponses[a.ID] = responsesForArtist
	}

	return idToResponses, nil
}

func (u *User) getArtistsAlbumsForArtist(artist model.Artist) ([]*model.GetArtistsAlbumsResponse, error) {
	var responses []*model.GetArtistsAlbumsResponse

	offset := 0
	albumLen := 0
	total := 1
	batchSize := 50
	for albumLen < total {
		inputs := model.NewRequestInput(&artist.ID, nil, nil, &offset, &batchSize, nil)
		response, err := u.api.DoRequest(u.artistConverter.BuildGetArtistsAlbumsRequest, inputs)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		decodedResponse, err := u.artistConverter.DecodeGetArtistsAlbumsResponse(*response)
		if err != nil {
			return nil, err
		}

		// set after to the last artist in the list
		offset = decodedResponse.Offset

		// set total number of artists
		total = decodedResponse.Total

		// set let of processed artists
		albumLen += len(decodedResponse.Albums)

		// add response to responses
		responses = append(responses, decodedResponse)
	}

	return responses, nil
}
