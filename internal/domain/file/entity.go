package file

import "time"

// FileMeta описание метаданных файла
type FileMeta struct {
	Name      string    // Имя файла
	CreatedAt time.Time // Дата создания файла
	UpdatedAt time.Time // Дата последнего обновления файла
}
