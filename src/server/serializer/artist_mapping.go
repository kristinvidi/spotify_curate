package serializer

import (
	"src/domain/model"
	pb "src/server/proto"

	"go.openly.dev/pointy"
)

func SerializeGetUnmappedArtistsForUserResponse(success bool, artists []model.Artist) *pb.GetUnmappedArtistsForUserResponse {
	return &pb.GetUnmappedArtistsForUserResponse{
		General: SerializeGeneral(success, nil),
		Artists: serverArtistsFromDomainArtists(artists),
	}
}

func serverArtistsFromDomainArtists(domainArtists []model.Artist) []*pb.Artist {
	artists := make([]*pb.Artist, len(domainArtists))
	for i, a := range domainArtists {
		artists[i] = serverArtistFromArtist(a)
	}

	return artists
}

func serverArtistFromArtist(artist model.Artist) *pb.Artist {
	return &pb.Artist{
		Id:        string(artist.ID),
		Uri:       string(artist.URI),
		Name:      string(artist.Name),
		CreatedAt: TimeToPbTimestamp(&artist.CreatedAt),
	}
}

func DeserializeCreateGenreToArtistsMappingsRequest(request *pb.CreateGenreToArtistsMappingsRequest) []model.GenreToArtistsMapping {
	mappings := make([]model.GenreToArtistsMapping, len(request.GenreToArtistsMappings))
	for i, mapping := range request.GenreToArtistsMappings {
		mappings[i] = model.GenreToArtistsMapping{
			Genre:       mapping.Genre,
			ArtistNames: mapping.ArtistNames,
		}
	}

	return mappings
}

func SerializeCreateGenreToArtistsMappingsResponse(success bool, unfollowedArtists []model.GenreToArtistsMapping) *pb.CreateGenreToArtistsMappingsResponse {
	var failureDetails *string
	if len(unfollowedArtists) > 0 {
		failureDetails = pointy.String("some artist mappings were not generated as the artists are not followed by the user")
	}

	return &pb.CreateGenreToArtistsMappingsResponse{
		General:                      SerializeGeneral(success, failureDetails),
		FailedGenreToArtistsMappings: SerializeGenreToArtistsMappings(unfollowedArtists),
	}
}

func SerializeGenreToArtistsMappings(mappings []model.GenreToArtistsMapping) []*pb.GenreToArtistsMapping {
	pbMappings := make([]*pb.GenreToArtistsMapping, len(mappings))
	for i, m := range mappings {
		pbMappings[i] = &pb.GenreToArtistsMapping{
			Genre:       m.Genre,
			ArtistNames: m.ArtistNames,
		}
	}

	return pbMappings
}
