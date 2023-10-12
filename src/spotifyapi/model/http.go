package model

import "fmt"

type RequestInput struct {
	id        *ID
	after     *string
	offset    *int
	batchSize *int
}

func NewRequestInput(id *ID, after *string, offset, batchSize *int) *RequestInput {
	return &RequestInput{id: id, after: after, offset: offset, batchSize: batchSize}
}

func (r *RequestInput) IDString() *string {
	if r.id == nil {
		return nil
	}

	str := r.id.String()
	return &str
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
