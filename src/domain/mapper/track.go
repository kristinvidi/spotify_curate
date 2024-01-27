package mapper

import (
	"strings"

	api "src/spotifyapi/model"
)

func TrackAPIURIsFromGetAlbumTracksResponses(responses []*api.GetAlbumTracksResponse) []api.URI {
	var trackIDs []api.URI

	for _, r := range responses {
		if r == nil {
			continue
		}
		for _, t := range r.Tracks {
			if !trackHasMixed(t) {
				trackIDs = append(trackIDs, t.URI)
			}
		}
	}

	return trackIDs
}

func trackHasMixed(t api.SimplifiedTrack) bool {
	return strings.Contains(strings.ToLower(t.Name), " - mixed")
}
