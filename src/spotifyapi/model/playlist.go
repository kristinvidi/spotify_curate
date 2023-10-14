package model

type CreatePlaylistInputs struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	Name          string `json:"name"`
	Public        bool   `json:"public"`
}

type CreatePlaylistResponse struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ID            ID     `json:"id"`
	Name          string `json:"name"`
	Public        bool   `json:"public"`
}

func NewCreatePlaylistInputs(name string, public, collaborative bool, description string) *CreatePlaylistInputs {
	return &CreatePlaylistInputs{
		Name:          name,
		Public:        public,
		Collaborative: collaborative,
		Description:   description,
	}
}

type AddTracksToPlaylistInputs struct {
	TrackURIs []URI `json:"uris"`
}

type AddTracksToPlaylistResponse struct {
	SnapshotID string `json:"snapshot_id"`
}
