package server

import (
	"src/domain"
	"src/domain/mapper"
	pb "src/server/proto"
	"src/server/serializer"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func (g *GrpcServer) UpdateUserData(ctx context.Context, request *pb.UpdateUserDataRequest) (*pb.UpdateUserDataResponse, error) {
	api := API_UPDATE_USER_DATA
	g.logAPICallStart(api)

	updater := domain.NewUserUpdater(g.config, g.logger)

	err := updater.UpdateUserData()
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeUpdateUserData(false), err
	}

	g.logAPICallSuccess(api)

	return serializer.SerializeUpdateUserData(true), nil
}

func (g *GrpcServer) GetUnmappedArtistsForUser(ctx context.Context, request *pb.GetUnmappedArtistsForUserRequest) (*pb.GetUnmappedArtistsForUserResponse, error) {
	api := API_GET_UNMAPPED_ARTISTS_FOR_USER
	g.logAPICallStart(api)

	updater := domain.NewUserUpdater(g.config, g.logger)

	artists, err := updater.GetUnmappedArtistsForUser()
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeGetUnmappedArtistsForUser(false, nil), err
	}

	g.logAPICallSuccess(api)

	return serializer.SerializeGetUnmappedArtistsForUser(
		true,
		mapper.ServerArtistsFromDomainArtists(artists),
	), nil
}

func (g *GrpcServer) CreatePlaylistRecentInGenre(ctx context.Context, request *pb.CreatePlaylistRecentInGenreRequest) (*pb.CreatePlaylistRecentInGenreResponse, error) {
	api := API_CREATE_PLAYLIST_RECENT_IN_GENRE
	g.logAPICallStart(api)

	genre, err := serializer.DeserializeCreatePlaylistRecentInGenreRequest(request)
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeCreatePlaylistRecentInGenreResponse(false), err
	}

	creator := domain.NewPlaylistCreator(g.config)

	generated, err := creator.CreateRecentInGenre(genre)
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeCreatePlaylistRecentInGenreResponse(false), err
	}

	if !generated {
		g.logger.Info("no new content to add to playlist, skipping generating recent in genre playlist", zap.String("genre", genre))
	} else {
		g.logAPICallSuccess(api)
	}

	return serializer.SerializeCreatePlaylistRecentInGenreResponse(true), nil
}

func (g *GrpcServer) CreatePlaylistRecentInGenreAll(ctx context.Context, request *pb.CreatePlaylistRecentInGenreAllRequest) (*pb.CreatePlaylistRecentInGenreAllResponse, error) {
	api := API_CREATE_PLAYLIST_RECENT_IN_GENRE_ALL
	g.logAPICallStart(api)

	creator := domain.NewPlaylistCreator(g.config)

	count, err := creator.CreateRecentInGenreAll()
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeCreatePlaylistRecentInGenreAllResponse(false), err
	}

	g.logger.Info("created playlists for all mapped genres", zap.Int("playlistCount", count))

	g.logAPICallSuccess(api)

	return serializer.SerializeCreatePlaylistRecentInGenreAllResponse(true), nil
}
