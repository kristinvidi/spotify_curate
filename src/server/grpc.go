package server

import (
	"src/config"
	"src/domain"
	"src/domain/mapper"
	pb "src/server/proto"
	"src/server/serializer"

	"go.uber.org/zap"
)

type apiEndpoint string

const (
	API_UPDATE_USER_DATA                    apiEndpoint = "update_user_data"
	API_GET_UNMAPPED_ARTISTS_FOR_USER       apiEndpoint = "get_unmapped_artists_for_user"
	API_CREATE_PLAYLIST_RECENT_IN_GENRE     apiEndpoint = "create_playlist_recent_in_genre"
	API_CREATE_PLAYLIST_RECENT_IN_GENRE_ALL apiEndpoint = "create_playlist_recent_in_genre_all"
)

// TO-DO: Add grpc server logic, currently it's just an abstracted layer to facilitate building out the grpc code
type GrpcServer struct {
	config *config.Config
	logger *zap.Logger
}

func NewGrpcServer(config *config.Config, logger *zap.Logger) *GrpcServer {
	return &GrpcServer{
		config: config,
		logger: logger,
	}
}

func (g *GrpcServer) logAPICall(apiEndpoint apiEndpoint) {
	g.logger.Info("calling api", zap.String("api", string(apiEndpoint)))
}

func (g *GrpcServer) logError(apiEndpoint apiEndpoint, err error) {
	g.logger.Error("error encountered during api call", zap.String("api", string(apiEndpoint)), zap.String("error", err.Error()))
}

func (g *GrpcServer) UpdateUserData(request *pb.UpdateUserDataRequest) (*pb.UpdateUserDataResponse, error) {
	api := API_UPDATE_USER_DATA
	g.logAPICall(api)

	updater := domain.NewUserUpdater(g.config, g.logger)

	err := updater.UpdateUserData()
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeUpdateUserData(false), err
	}

	g.logger.Info("successfully updated user data")

	return serializer.SerializeUpdateUserData(true), nil
}

func (g *GrpcServer) GetUnmappedArtistsForUser(request *pb.GetUnmappedArtistsForUserRequest) (*pb.GetUnmappedArtistsForUserResponse, error) {
	api := API_GET_UNMAPPED_ARTISTS_FOR_USER
	g.logAPICall(api)

	updater := domain.NewUserUpdater(g.config, g.logger)

	artists, err := updater.GetUnmappedArtistsForUser()
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeGetUnmappedArtistsForUser(false, nil), err
	}

	g.logger.Info("successfully fetched unmapped artists for user")

	return serializer.SerializeGetUnmappedArtistsForUser(
		true,
		mapper.ServerArtistsFromDomainArtists(artists),
	), nil
}

func (g *GrpcServer) CreatePlaylistRecentInGenre(request *pb.CreatePlaylistRecentInGenreRequest) (*pb.CreatePlaylistRecentInGenreResponse, error) {
	api := API_CREATE_PLAYLIST_RECENT_IN_GENRE
	g.logAPICall(api)

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
		g.logger.Info("successfully created recent in genre playlist", zap.String("genre", genre))
	}

	return serializer.SerializeCreatePlaylistRecentInGenreResponse(true), nil
}

func (g *GrpcServer) CreatePlaylistRecentInGenreAll(request *pb.CreatePlaylistRecentInGenreAllRequest) (*pb.CreatePlaylistRecentInGenreAllResponse, error) {
	api := API_CREATE_PLAYLIST_RECENT_IN_GENRE_ALL
	g.logAPICall(api)

	creator := domain.NewPlaylistCreator(g.config)

	count, err := creator.CreateRecentInGenreAll()
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeCreatePlaylistRecentInGenreAllResponse(false), err
	}

	g.logger.Info("successfully created recent in genre playlists for mapped genres", zap.Int("playlistCount", count))

	return serializer.SerializeCreatePlaylistRecentInGenreAllResponse(true), nil
}
