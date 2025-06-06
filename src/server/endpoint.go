package server

import (
	"src/domain"
	pb "src/server/proto"
	"src/server/serializer"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type apiEndpoint string

const (
	API_AUTHENTICATE_USER                   apiEndpoint = "authenticate_user"
	API_UPDATE_USER_DATA                    apiEndpoint = "update_user_data"
	API_GET_UNMAPPED_ARTISTS_FOR_USER       apiEndpoint = "get_unmapped_artists_for_user"
	API_CREATE_PLAYLIST_RECENT_IN_GENRE     apiEndpoint = "create_playlist_recent_in_genre"
	API_CREATE_PLAYLIST_RECENT_IN_GENRE_ALL apiEndpoint = "create_playlist_recent_in_genre_all"
	API_CREATE_ARTIST_GENRE_MAPPING         apiEndpoint = "create_artist_genre_mapping"
	API_CREATE_LABELS_FOR_USER              apiEndpoint = "create_labels_for_user"
)

func (g *GrpcServer) AuthenticateUser(ctx context.Context, request *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	api := API_AUTHENTICATE_USER
	g.logAPICallStart(api)

	manager := domain.NewUserManager(g.config, g.logger)

	userID, err := manager.AuthenticateUser()
	if err != nil {
		g.logError(api, err)
		return serializer.SerializeAuthenticateUserResponse(false, nil), err
	}

	g.logAPICallSuccess(api)

	return serializer.SerializeAuthenticateUserResponse(true, userID), nil
}

func (g *GrpcServer) UpdateUserData(ctx context.Context, request *pb.UpdateUserDataRequest) (*pb.UpdateUserDataResponse, error) {
	api := API_UPDATE_USER_DATA
	g.logAPICallStart(api)

	updater := domain.NewUserManager(g.config, g.logger)

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

	userID := request.GetUserSpotifyId()
	updater := domain.NewUserManager(g.config, g.logger)

	artists, err := updater.GetUnmappedArtistsForUser(userID)
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeGetUnmappedArtistsForUserResponse(false, nil), err
	}

	g.logAPICallSuccess(api)

	return serializer.SerializeGetUnmappedArtistsForUserResponse(true, artists), nil
}

func (g *GrpcServer) CreateGenreToArtistsMappings(ctx context.Context, request *pb.CreateGenreToArtistsMappingsRequest) (*pb.CreateGenreToArtistsMappingsResponse, error) {
	api := API_CREATE_ARTIST_GENRE_MAPPING
	g.logAPICallStart(api)

	userID := request.GetUserSpotifyId()
	mappings := serializer.DeserializeCreateGenreToArtistsMappingsRequest(request)

	updater := domain.NewUserManager(g.config, g.logger)

	unfollowedArtists, err := updater.CreateArtistToGenreMappingForUser(userID, mappings)
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeCreateGenreToArtistsMappingsResponse(false, nil), err
	}

	g.logAPICallSuccess(api)

	return serializer.SerializeCreateGenreToArtistsMappingsResponse(true, unfollowedArtists), nil
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

	err = creator.CreateRecentInGenre(genre)
	if err != nil {
		g.logError(api, err)

		return serializer.SerializeCreatePlaylistRecentInGenreResponse(false), err
	}

	g.logAPICallSuccess(api)

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

func (g *GrpcServer) CreateLabelsForUser(ctx context.Context, request *pb.CreateLabelsForUserRequest) (*pb.CreateLabelsForUserResponse, error) {
	api := API_CREATE_LABELS_FOR_USER
	g.logAPICallStart(api)

	userID := request.GetUserSpotifyId()
	labels := request.GetLabels()

	updater := domain.NewUserManager(g.config, g.logger)

	failedLabels, err := updater.CreateLabelsForUser(userID, labels)
	if err != nil {
		g.logError(api, err)
		return serializer.SerializeCreateLabelsForUserResponse(false, failedLabels), err
	}

	g.logAPICallSuccess(api)

	return serializer.SerializeCreateLabelsForUserResponse(true, failedLabels), nil
}
