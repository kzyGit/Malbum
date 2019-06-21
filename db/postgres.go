package db

import (
	"database/sql"
	"../schema"
	_ "github.com/lib/pq"
)

type Postgres struct{
	DB *sql.DB
}

func ConnectPostgres() (*Postgres, error) {
	connStr := "postgresql://postgres@localhost:5432/malbum?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
		panic(err)
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return &Postgres{db}, nil
}

func (p *Postgres) Close(){}

func (p *Postgres) Insert(song *schema.Album) (int, error) {

	query := `INSERT INTO malbum (id, title, author, dateadded)
	VALUES (nextval('song_id'), $1, $2, $3) RETURNING *;`
	
	rows, err := p.DB.Query(query, song.Title, song.Artist, song.DateAdded)
    if err != nil {
        return -1, err
    }

    var id int
    for rows.Next() {
        if err := rows.Scan(&id); err != nil {
            return -1, err
        }
    }

    return id, nil
}

func (p *Postgres) Delete(id int) error {
    query := `
        DELETE FROM malbum
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
        FROM malbum
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

    return songsList, nil
}