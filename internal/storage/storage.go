package storage

import "github.com/istrel/storage/internal/file"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

//Загрузка файла в хранилище.
func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return file.NewFile(filename, blob)
	// if err != nil {
	// 	return nil, err
	// }

	// return newFile
}
