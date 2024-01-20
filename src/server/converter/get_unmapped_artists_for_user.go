package converter

import pb "src/server/proto"

func SerializeGetUnmappedArtistsForUser(success bool, artists []*pb.Artist) *pb.GetUnmappedArtistsForUserResponse {
	return &pb.GetUnmappedArtistsForUserResponse{
		General: SerializeGeneral(success),
		Artists: artists,
	}
}
