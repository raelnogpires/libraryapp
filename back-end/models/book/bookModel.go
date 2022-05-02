package book

type UseCase interface {
	GetAll() ([]*Book, error)
	GetById(Id int64) (*Book, error)
	Create(b *Book) error
	Update(b *Book) error
	Delete(Id int64) error
}
