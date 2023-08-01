package models

import (
	"database/sql"
	"errors"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// GetAlbumByID récupère un album depuis la base de données en utilisant son ID.
func GetAlbumByID(db *sql.DB, id string) (*Album, error) {
	row := db.QueryRow("SELECT id, title, artist, price FROM albums WHERE id = $1", id)
	album := &Album{}
	err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("album not found")
		}
		return nil, err
	}
	return album, nil
}

// Save enregistre un nouvel album dans la base de données.
func (album *Album) Save(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO albums (id, title, artist, price) VALUES ($1, $2, $3, $4)",
		album.ID, album.Title, album.Artist, album.Price)
	if err != nil {
		return err
	}
	return nil
}

// GetAllAlbums récupère tous les albums depuis la base de données.
func GetAllAlbums(db *sql.DB) ([]*Album, error) {
	rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albums := []*Album{}
	for rows.Next() {
		album := &Album{}
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}