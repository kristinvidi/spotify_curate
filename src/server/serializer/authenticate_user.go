package serializer

import pb "src/server/proto"

func SerializeAuthenticateUserResponse(success bool, userID *string) *pb.AuthenticateUserResponse {
	var id string
	if userID != nil {
		id = *userID
	}
	return &pb.AuthenticateUserResponse{
		UserSpotifyId: id,
		General:       SerializeGeneral(success, nil),
	}
}
