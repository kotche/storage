package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/istrel/storage/internal/file"
)

//в мапе храним файлы и их id
type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

//Загрузка файла в хранилище.
func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(fileId uuid.UUID) (*file.File, error) {
	foundFile, exists := s.files[fileId]

	if !exists {
		return nil, fmt.Errorf("file %v not found", fileId)
	}

	return foundFile, nil
}
