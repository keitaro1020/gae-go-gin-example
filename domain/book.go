package domain

type Book struct {
	ID     string `json:"id" datastore:"-" boom:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int64  `json:"price"`
}

func NewBook(id string) *Book {
	b := &Book{
		ID: id,
	}
	return b
}
