package datastore

import (
	"fmt"
	"os"
)

type TextDB struct{}

func (t *TextDB) buildAbsoluteFilepath(filename string) (string, error) {
	// relativeFilepath := fmt.Sprintf("../data_store/%s", filename)

	// return filepath.Abs(relativeFilepath)

	path := fmt.Sprintf("/Users/kiki.vidi/Apps/spotify_curate/api/data_store/%s", filename)

	return path, nil
}

func (t *TextDB) PrintAThing() {
	fmt.Println("Im a thing!")
}

// func (t *TextDB) GetEntry(key string) interface{} {

// }

func (t *TextDB) WriteEntry() error {

	return nil
}

func (t *TextDB) WriteAllEntries(filename string, data []byte) error {
	path, err := t.buildAbsoluteFilepath(filename)

	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (t *TextDB) ReadAllEntries(filename string) ([]byte, error) {
	path, err := t.buildAbsoluteFilepath(filename)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}
