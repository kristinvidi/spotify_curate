package datastore

type DataStore interface {
	WriteAllEntries(filename string, data []byte) error
}
