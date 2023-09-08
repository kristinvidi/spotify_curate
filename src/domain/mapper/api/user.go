package mapperapi

import (
	"src/domain/model"
	api "src/spotifyapi/model"
)

func UserFromCurrentUsersProfileResponse(response *api.GetCurrentUsersProfileResponse) *model.User {
	if response == nil {
		return nil
	}

	return &model.User{
		DisplayName: model.Name(response.DisplayName),
		Email:       model.Email(response.Email),
		ID:          model.ID(response.ID),
		URI:         model.URI(response.URI),
		Country:     model.CountryCode(response.Country),
	}
}
