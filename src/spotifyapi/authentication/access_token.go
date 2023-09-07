package authentication

import (
	"os"
	"path/filepath"

	"src/spotifyapi/model"
)

const filename = "access_token.txt"

type AccessTokenStorage struct{}

func NewAccessTokenStorage() *AccessTokenStorage {
	return &AccessTokenStorage{}
}

func (a *AccessTokenStorage) GetFromFile() (*model.AccessToken, error) {
	path, err := a.absolutePath(filename)
	if err != nil {
		return nil, err
	}

	if !a.fileExists(path) {
		return nil, nil
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	token := model.AccessToken(contents)

	return &token, nil
}

func (a *AccessTokenStorage) WriteToFile(token model.AccessToken) error {
	path, err := a.absolutePath(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(string(token))
	if err != nil {
		return err
	}

	return nil
}

func (a *AccessTokenStorage) fileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err == nil {
		return true

	} else {
		return false
	}
}

func (a *AccessTokenStorage) absolutePath(filename string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	authPath := filepath.Join(cwd, "spotifyapi", "authentication", filename)

	return authPath, nil
}
