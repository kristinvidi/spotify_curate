package mapper

import api "src/spotifyapi/model"

func TrackAPIURIsFromGetAlbumTracksResponses(responses []*api.GetAlbumTracksResponse) []api.URI {
	var trackIDs []api.URI

	for _, r := range responses {
		if r == nil {
			continue
		}
		for _, t := range r.Tracks {
			trackIDs = append(trackIDs, t.URI)
		}
	}

	return trackIDs
}
