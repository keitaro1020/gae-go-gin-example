package domain

type Book struct {
	ID     string   `json:"-" datastore:"-" boom:"id"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Price  int64    `json:"price"`
	Ref    *BookRef `json:"-" datastore:"-"`
}

type BookRef struct {
	ID     *string `json:"id"`
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Price  *int64  `json:"price"`
}

func NewBook(id string) *Book {
	b := &Book{
		ID:  id,
		Ref: &BookRef{},
	}

	b.Ref.ID = &b.ID
	b.Ref.Title = &b.Title
	b.Ref.Author = &b.Author
	b.Ref.Price = &b.Price

	return b
}
