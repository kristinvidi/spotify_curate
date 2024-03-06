package model

type SimplifiedTrack struct {
	Name string
	ID   ID
	URI  URI
}

type SimplifiedTracks []SimplifiedTrack

func (s SimplifiedTracks) URIs() []URI {
	var uris []URI
	for _, t := range s {
		uris = append(uris, t.URI)
	}
	return uris
}

func (s SimplifiedTracks) IDs() []ID {
	var ids []ID
	for _, t := range s {
		ids = append(ids, t.ID)
	}
	return ids
}
