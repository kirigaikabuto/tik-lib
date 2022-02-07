package tik_lib

type FileStore interface {
	Create(file *File) (*File, error)
	Update(file *FileUpdate) (*File, error)
	Get(id string) (*File, error)
	Delete(id string) error
	List() ([]File, error)
}
