package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

//Конструктор файла. Если возвращать File, придется описывать File{}.
//При возврате указателя *File при ошибке достаточно вернуть nil
//При ошибке стараться возвращать nil
func NewFile(fileName string, blob []byte) (*File, error) {
	fileId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   fileId,
		Name: fileName,
		Data: blob,
	}, nil
}

//переопределим вывод данных о файле
func (f *File) String() string {
	return fmt.Sprintf("File(%s, %v)", f.Name, f.ID)
}
