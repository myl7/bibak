package db

import (
	"errors"
	"github.com/myl7/bibak/internal/model"
)

func SelectFavItem(bid string) (*model.FavItem, error) {
	rows, err := DB.Queryx(
		"SELECT bid, title, description, add_fav_time, link, author FROM fav_items WHERE bid = $1",
		bid,
	)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New("no result")
	}

	var f model.FavItem
	err = rows.StructScan(&f)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func InsertFavItem(f *model.FavItem) error {
	_, err := DB.NamedExec(
		"INSERT INTO fav_items (bid, title, description, add_fav_time, link, author) VALUES (:bid, :title, :description, :add_fav_time, :link, :author)",
		*f,
	)
	if err != nil {
		return err
	}
	return nil
}
