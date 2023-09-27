package api

import (
	"src/spotifyapi/convert"
	"src/spotifyapi/model"

	"go.openly.dev/pointy"
)

type User struct {
	api       *API
	converter *convert.User
}

func NewUser(api *API, converter *convert.User) *User {
	return &User{api: api}
}

func (u *User) GetCurrentUsersProfile() (*model.GetCurrentUsersProfileResponse, error) {
	response, err := u.api.DoRequest(u.converter.BuildCurrentUsersProfileRequest, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decodedResponse, err := u.converter.DecodeCurrentUsersProfileResponse(*response)
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
		inputs := model.NewRequestInput(after, &batchSize)
		response, err := u.api.DoRequest(u.converter.BuildGetFollowedArtistsRequest, inputs)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		decodedResponse, err := u.converter.DecodeGetFollowedArtistsResponse(*response)
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
