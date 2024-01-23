package serializer

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type deserializeError struct {
	message string
}

func (s *deserializeError) Error() error {
	return status.Errorf(codes.InvalidArgument, s.message)
}
