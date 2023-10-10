package model

import "fmt"

type RequestInput struct {
	spotifyID *string
	after     *string
	offset    *int
	batchSize *int
}

func NewRequestInput(spotifyID, after *string, offset, batchSize *int) *RequestInput {
	return &RequestInput{spotifyID: spotifyID, after: after, offset: offset, batchSize: batchSize}
}

func (r *RequestInput) SpotifyID() *string {
	return r.spotifyID
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
