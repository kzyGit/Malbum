package db

import (
	"database/sql"

	"fmt"

	"../schema"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func ConnectPostgres() (*Postgres, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=malbum sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

func (p *Postgres) Close() {}

func (p *Postgres) Insert(song *schema.Album) (string, schema.Album, error) {

	query := `INSERT INTO album (id, title, artist, dateadded)
	VALUES (nextval('song_id'), $1, $2, $3) RETURNING *;`

	rows, err := p.DB.Query(query, song.Title, song.Artist, song.DateAdded)
    var newsong schema.Album
    
    if err != nil {
		return "Error", newsong, err
	}

	// var id int
	// for rows.Next() {
	// 	if err := rows.Scan(&id); err != nil {
	// 		return -1, err
	// 	}
	// }

	
	for rows.Next() {

		if err := rows.Scan(&newsong.ID, &newsong.Title, &newsong.Artist, &newsong.DateAdded); err != nil {
			return "Error", newsong, err
		}

	}
	return "Song added successfully", newsong, nil
}

func (p *Postgres) Delete(id int) error {
	query := `
        DELETE FROM album
        WHERE id = $1;
    `

	if _, err := p.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetAll() ([]schema.Album, error) {
	query := `
        SELECT *
        FROM album
        ORDER BY id;
    `

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var songsList []schema.Album
	for rows.Next() {
		var t schema.Album
		if err := rows.Scan(&t.ID, &t.Title, &t.Artist, &t.DateAdded); err != nil {
			return nil, err
		}
		songsList = append(songsList, t)
	}
	fmt.Println("all **", songsList)

	return songsList, nil
}

func (p *Postgres) GetOne(id int) (schema.Album, error) {
	query := `
        SELECT *
        FROM album
        WHERE id = $1;
    `

	rows, err := p.DB.Query(query, id)
	var t schema.Album
	if err != nil {
		return t, nil
	}

	for rows.Next() {
		if err := rows.Scan(&t.ID, &t.Title, &t.Artist, &t.DateAdded); err != nil {
			return t, nil
		}
	}
	return t, nil
}
