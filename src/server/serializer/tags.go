package serializer

import pb "src/server/proto"

func SerializeGetArtistTagsForUserResponse(success bool, tags []string) *pb.GetArtistTagsForUserResponse {
	return &pb.GetArtistTagsForUserResponse{
		General: SerializeGeneral(success, nil),
		Tags:    tags,
	}
}
