package config

import (
	"fmt"
	"os"
	"path/filepath"

	apptype "spotify_app/api/pkg/app_type"
	"spotify_app/api/pkg/constants"
)

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) GetAccessToken() (*apptype.AccessToken, error) {
	tokenString, err := m.getEnvironmentVariable(constants.EnvAccessToken)
	if err != nil {
		return nil, err
	}

	token := apptype.AccessToken(tokenString)

	return &token, nil
}

func (m *Manager) WriteAccessTokenToFile(token apptype.AccessToken) error {
	return m.writeEnvironmentVariableToFile(constants.EnvAccessToken, string(token))
}

func (m *Manager) GetUserCountryCode() (*apptype.UserCountryCode, error) {
	codeString, err := m.getEnvironmentVariable(constants.EnvUserCountryCode)
	if err != nil {
		return nil, err
	}

	code := apptype.UserCountryCode(codeString)

	return &code, nil
}

func (m *Manager) WriteUserCountryCodeToFile(code apptype.UserCountryCode) error {
	return m.writeEnvironmentVariableToFile(constants.EnvUserCountryCode, string(code))
}

func (m *Manager) GetUserID() (*apptype.UserID, error) {
	codeString, err := m.getEnvironmentVariable(constants.EnvUserID)
	if err != nil {
		return nil, err
	}

	code := apptype.UserID(codeString)

	return &code, nil
}

func (m *Manager) WriteUserIDToFile(id apptype.UserID) error {
	return m.writeEnvironmentVariableToFile(constants.EnvUserID, string(id))
}

func (m *Manager) buildAbsoluteFilepath(name constants.EnvironmentVariable) (string, error) {
	relativeFilepath := fmt.Sprintf("../config/environment_files/%s", name.Path())

	return filepath.Abs(relativeFilepath)
}

func (m *Manager) getEnvironmentVariable(name constants.EnvironmentVariable) (string, error) {
	path, err := m.buildAbsoluteFilepath(name)
	if err != nil {
		return "", err
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func (m *Manager) writeEnvironmentVariableToFile(name constants.EnvironmentVariable, value string) error {
	path, err := m.buildAbsoluteFilepath(name)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(string(value))
	if err != nil {
		return err
	}

	return nil
}
