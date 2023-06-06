package model

type Track struct {
	Album            Album    `json:"album"`
	Artists          Artists  `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	DiscNumber       int      `json:"disc_number"`
	DurationMS       int      `json:"duration_ms"`
	Explicit         bool     `json:"explicit"`
	ExternalIDs      struct {
		ISRC string `json:"isrc"`
		EAN  string `json:"ean"`
		UPC  string `json:"upc"`
	} `json:"external_ids"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	HREF       string `json:"href"`
	ID         string `json:"id"`
	IsPlayable bool   `json:"is_playable"`
	// LinkedFrom
	Restrictions struct {
		Reason string `json:"reason"`
	}
	Name        string `json:"name"`
	Popularity  int    `json:"popularity"`
	PreviewURL  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string `json:"track"`
	URI         string `json:"uri"`
	IsLocal     bool   `json:"is_local"`
}

type Tracks []Track

func (t Tracks) URIs() []string {
	var ids []string
	for _, track := range t {
		ids = append(ids, track.URI)
	}

	return ids
}

// func (t Tracks) GetFirstGenreFromRelatedArtist() apptype.Genre {
// 	mapGenreToCount := make(map[Genre]int)
//
// 	for _, track := range t {
// 		for _, artist := track.Artists {
// 			if artist
// 		}
// 	}
// }
//
// func getGenreFromArtists(artists model.Artists) apptype.
