package mapper

import (
	"strings"

	"src/domain/model"
	api "src/spotifyapi/model"
)

func SimplifiedTrackFromAPISimplifiedTrack(t api.SimplifiedTrack) model.SimplifiedTrack {
	return model.SimplifiedTrack{
		Name: t.Name,
		ID:   model.ID(t.ID),
		URI:  model.URI(t.URI),
	}
}

func SimplifiedTracksFromGetAlbumTracksResponses(responses []*api.GetAlbumTracksResponse) model.SimplifiedTracks {
	var tracks model.SimplifiedTracks

	for _, r := range responses {
		if r == nil {
			continue
		}
		for _, t := range r.Tracks {
			if !trackHasMixed(t) {
				tracks = append(tracks, SimplifiedTrackFromAPISimplifiedTrack(t))
			}
		}
	}

	return tracks
}

func trackHasMixed(t api.SimplifiedTrack) bool {
	return strings.Contains(strings.ToLower(t.Name), " - mixed")
}
