package serializer

import pb "src/server/proto"

func DeserializeCreatePlaylistRecentInGenreRequest(request *pb.CreatePlaylistRecentInGenreRequest) (string, error) {
	if request == nil {
		return "", (&deserializeError{message: "request is nil"}).Error()
	}

	return request.Genre, nil
}

func SerializeCreatePlaylistRecentInGenreResponse(success bool) *pb.CreatePlaylistRecentInGenreResponse {
	return &pb.CreatePlaylistRecentInGenreResponse{
		General: SerializeGeneral(success, nil),
	}
}

func SerializeCreatePlaylistRecentInGenreAllResponse(success bool) *pb.CreatePlaylistRecentInGenreAllResponse {
	return &pb.CreatePlaylistRecentInGenreAllResponse{
		General: SerializeGeneral(success, nil),
	}
}
