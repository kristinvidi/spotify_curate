package model

import apptype "spotify_app/api/pkg/app_type"

type CreatePlaylistRequest struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	Name          string `json:"name"`
	Public        bool   `json:"public"`
}

func NewCreatePlaylistRequest(name string, public, collaborative bool, description string) *CreatePlaylistRequest {
	return &CreatePlaylistRequest{
		Name:          name,
		Public:        public,
		Collaborative: collaborative,
		Description:   description,
	}
}

type CreatePlaylistResponse struct {
	Collaborative bool               `json:"collaborative"`
	Description   string             `json:"description"`
	ID            apptype.PlaylistID `json:"id"`
	Name          string             `json:"name"`
	Public        bool               `json:"public"`
}

type AddItemsToPlaylistRequest struct {
	URIs     []string `json:"uris"`
	Position int      `json:"position"`
}

func NewAddItemsToPlaylistRequest(ids []string) *AddItemsToPlaylistRequest {
	return &AddItemsToPlaylistRequest{
		URIs:     ids,
		Position: 0,
	}
}
