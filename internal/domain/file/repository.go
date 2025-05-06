package file

import (
	"context"
	"io"
)

// FileRepository описывает абстрактное файловое хранилище
type FileRepository interface {
	Save(ctx context.Context, name string, content io.Reader) error // Сохранить файл
	Load(ctx context.Context, name string) (io.ReadCloser, error)   // Загрузить файл
	List(ctx context.Context) ([]FileMeta, error)                   // Получить список файлов
}
