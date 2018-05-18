package repository

import (
	"context"
	"go.mercari.io/datastore/boom"
	"go.mercari.io/datastore"
)

func BoomFromContext(ctx context.Context) (*boom.Boom, error) {
	return boom.FromContext(ctx)
}

func GetKeys(ctx context.Context, q datastore.Query) ([]datastore.Key, error) {
	bm, err := BoomFromContext(ctx)
	if err != nil {
		return nil, err
	}

	it := bm.Run(q)
	keys := []datastore.Key{}

	for {
		if key, err := it.Next(nil); err != nil {
			break
		} else {
			keys = append(keys, key)
		}
	}
	return keys, nil
}