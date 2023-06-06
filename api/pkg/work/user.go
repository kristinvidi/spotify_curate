package work

import (
	"spotify_app/api/config"
	apptype "spotify_app/api/pkg/app_type"
	httprequest "spotify_app/api/pkg/http_request"
)

func StoreUserInfo(httpRequest httprequest.HttpRequest, accessToken apptype.AccessToken, configManager *config.Manager) error {
	currentUser, err := httpRequest.GetCurrentUsersProfile(accessToken)
	if err != nil {
		return err
	}

	err = configManager.WriteUserCountryCodeToFile(currentUser.Country)
	if err != nil {
		return err
	}

	err = configManager.WriteUserIDToFile(currentUser.ID)
	if err != nil {
		return err
	}

	return nil
}
