package model

import "fmt"

type RequestInput struct {
	after     *string
	batchSize *int
}

func NewRequestInput(after *string, batchSize *int) *RequestInput {
	return &RequestInput{after: after, batchSize: batchSize}
}

func (r *RequestInput) After() *string {
	if r.after == nil {
		return nil
	}

	return r.after
}

func (r *RequestInput) BatchSize() *string {
	if r.batchSize == nil {
		return nil
	}

	batchSize := fmt.Sprintf("%d", *r.batchSize)

	return &batchSize
}
