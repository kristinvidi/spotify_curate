package mapper

import (
	db "src/db/model"
	"src/domain/model"
	pb "src/server/proto"
	"src/server/serializer"
)

func ArtistsFromDBArtists(dbArtists db.Artists) []model.Artist {
	artists := make([]model.Artist, len(dbArtists))
	for i, a := range dbArtists {
		artists[i] = ArtistFromDBArtist(a)
	}

	return artists
}

func ArtistFromDBArtist(dbArtist db.Artist) model.Artist {
	return model.Artist{
		ID:        model.ID(dbArtist.ID),
		URI:       model.URI(dbArtist.URI),
		Name:      dbArtist.Name,
		CreatedAt: dbArtist.CreatedAt,
	}
}

func ServerArtistsFromDomainArtists(domainArtists []model.Artist) []*pb.Artist {
	artists := make([]*pb.Artist, len(domainArtists))
	for i, a := range domainArtists {
		artists[i] = serverArtistFromDBArtist(a)
	}

	return artists
}

func serverArtistFromDBArtist(artist model.Artist) *pb.Artist {
	return &pb.Artist{
		Id:        string(artist.ID),
		Uri:       string(artist.URI),
		Name:      string(artist.Name),
		CreatedAt: serializer.TimeToPbTimestamp(&artist.CreatedAt),
	}
}
