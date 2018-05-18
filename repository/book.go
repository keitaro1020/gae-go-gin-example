package repository

import (
	"context"
	"github.com/keitaro1020/gae-go-gin-example/domain"
	"google.golang.org/appengine/log"
)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) GetBookByID(c context.Context, bId string) (*domain.Book, error) {
	bk := domain.Book{
		ID: bId,
	}

	bm, err := BoomFromContext(c)
	if err != nil {
		return nil, err
	}

	if err := bm.Get(&bk); err != nil {
		return nil, err
	}

	return &bk, nil
}

func (r *BookRepository) GetBooks(c context.Context) ([]*domain.Book, error) {

	bm, err := BoomFromContext(c)
	if err != nil {
		return nil, err
	}

	q := bm.Client.NewQuery(bm.Kind(&domain.Book{})).KeysOnly()
	keys, err := GetKeys(c, q)
	if err != nil {
		return nil, err
	}

	bs := []*domain.Book{}
	for _, v := range keys {
		bs = append(bs, domain.NewBook(v.Name()))
	}

	if err = bm.GetMulti(bs); err != nil {
		return nil, err
	}

	return bs, nil
}

func (r *BookRepository) PutBook(c context.Context, b *domain.Book) (*domain.Book, error) {
	bm, err := BoomFromContext(c)
	if err != nil {
		return nil, err
	}

	if _, err := bm.Put(b); err != nil {
		return nil, err
	}

	return b, nil
}
