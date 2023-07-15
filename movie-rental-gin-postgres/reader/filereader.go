package reader

import (
	"encoding/json"
	"errors"
	"io"
	"movieRentals/model"
	"os"
)

var ErrPathIsEmpty = errors.New("file path cannot be empty")
var ErrFileNotFound = errors.New("file not found")
var ErrFileCannotBeRead = errors.New("file cannot be read")
var ErrUnmarhsallingToJson = errors.New("cannot unmarshall json")

type FileReader struct {
	path string
}

func NewFileReader(path string) (*FileReader, error) {
	if path == "" {
		return nil, ErrPathIsEmpty
	}

	return &FileReader{path}, nil
}

func (fr *FileReader) GetMovies() (*model.Movies, error) {
	handle, err := os.Open(fr.path)

	if err != nil {
		return nil, ErrFileNotFound
	}

	defer handle.Close()
	return getMoviesFromReader(handle)
}

func getMoviesFromReader(handle io.Reader) (*model.Movies, error) {
	byteVal := make([]byte, 99999)
	n, err := handle.Read(byteVal)
	if err != nil {
		return nil, ErrFileCannotBeRead
	}

	var movies model.Movies
	b := byteVal[0:n]
	err = json.Unmarshal(b, &movies)
	if err != nil {
		return nil, ErrUnmarhsallingToJson
	}

	return &movies, nil
}
