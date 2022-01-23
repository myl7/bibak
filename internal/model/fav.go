package model

import "time"

type FavItem struct {
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	AddFavTime  time.Time `db:"add_fav_time" json:"addFavTime"`
	Link        string    `db:"link" json:"link"`
	Author      string    `db:"author" json:"author"`
	BID         string    `db:"bid" json:"bid"`
}
