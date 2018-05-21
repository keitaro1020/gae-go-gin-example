package repository

import (
	"context"
	"github.com/keitaro1020/gae-go-gin-example/domain"
)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) GetBookByID(c context.Context, bId string) (*domain.BookRef, error) {
	bm, err := BoomFromContext(c)
	if err != nil {
		return nil, err
	}

	bk := domain.NewBook(bId)

	if err := bm.Get(bk); err != nil {
		return nil, err
	}

	return bk.Ref, nil
}

func (r *BookRepository) GetBooks(c context.Context) ([]*domain.BookRef, error) {
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

	brs := []*domain.BookRef{}
	for _, v := range bs {
		brs = append(brs, v.Ref)
	}

	return brs, nil
}

func (r *BookRepository) PutBook(c context.Context, b *domain.Book) (*domain.BookRef, error) {
	bm, err := BoomFromContext(c)
	if err != nil {
		return nil, err
	}

	if _, err := bm.Put(b); err != nil {
		return nil, err
	}

	return b.Ref, nil
}
