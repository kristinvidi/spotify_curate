package model

import (
	"fmt"
)

type RequestInput struct {
	id                *ID
	uris              []URI
	after             *string
	offset            *int
	batchSize         *int
	newPlaylistInputs *CreatePlaylistInputs
}

func NewRequestInput(id *ID, uris []URI, after *string, offset, batchSize *int, newPlaylistInputs *CreatePlaylistInputs) *RequestInput {
	return &RequestInput{id: id, uris: uris, after: after, offset: offset, batchSize: batchSize, newPlaylistInputs: newPlaylistInputs}
}

func (r *RequestInput) IDString() *string {
	if r.id == nil {
		return nil
	}

	str := r.id.String()
	return &str
}

func (r *RequestInput) URIs() []URI {
	return r.uris
}

func (r *RequestInput) After() *string {
	return r.after
}

func (r *RequestInput) Offset() *string {
	if r.offset == nil {
		return nil
	}

	offset := fmt.Sprintf("%d", *r.offset)

	return &offset
}

func (r *RequestInput) BatchSize() *string {
	if r.batchSize == nil {
		return nil
	}

	batchSize := fmt.Sprintf("%d", *r.batchSize)

	return &batchSize
}

func (r *RequestInput) CreatePlaylistInputs() *CreatePlaylistInputs {
	return r.newPlaylistInputs
}

func (r *RequestInput) PlaylistName() *string {
	if r.newPlaylistInputs == nil {
		return nil
	}

	return &r.newPlaylistInputs.Name
}

func (r *RequestInput) PlaylistPublic() *bool {
	if r.newPlaylistInputs == nil {
		return nil
	}

	return &r.newPlaylistInputs.Public
}

func (r *RequestInput) PlaylistCollaborative() *bool {
	if r.newPlaylistInputs == nil {
		return nil
	}

	return &r.newPlaylistInputs.Collaborative
}

func (r *RequestInput) PlaylistDescription() *string {
	if r.newPlaylistInputs == nil {
		return nil
	}

	return &r.newPlaylistInputs.Description
}
