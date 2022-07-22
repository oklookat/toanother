package base

import (
	"database/sql"
	"errors"
)

type Album struct {
	ID     int64    `json:"id"`
	Title  string   `json:"title"`
	Artist []string `json:"artist"`
	// unix timestamp (ms).
	ReleaseDate int64 `json:"releaseDate"`
	TrackCount  int   `json:"trackCount"`
	Year        int   `json:"year"`
}

// delete all from table.
func (a *Album) DeleteAllFromTable(conn *sql.DB) (err error) {
	_, err = conn.Exec(`DELETE FROM liked_album`)
	return
}

// add new to table.
func (a *Album) AddToTable(conn *sql.DB) (id int64, err error) {
	if a.Artist == nil {
		err = errors.New("empty artist")
		return
	}
	var liked = Artist{}
	var artistConv = liked.sliceToNames(a.Artist)
	result, err := conn.Exec(`INSERT INTO 
	liked_album (title, artist, release_date, track_count, year) 
	VALUES ($1, $2, $3, $4, $5)`,
		a.Title, artistConv, a.ReleaseDate, a.TrackCount, a.Year)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}

// get albums list.
func (a *Album) GetAll(conn *sql.DB) (albums []*Album, err error) {
	rows, err := conn.Query("SELECT * FROM liked_album")
	if err != nil {
		return
	}
	defer rows.Close()
	albums = make([]*Album, 0)
	var liked = Artist{}
	for rows.Next() {
		var al = &Album{}
		var artistStr string
		err = rows.Scan(&al.ID, &al.Title, &artistStr, &al.ReleaseDate, &al.TrackCount, &al.Year)
		if err != nil {
			return
		}
		al.Artist = liked.namesToSlice(artistStr)
		albums = append(albums, al)
	}
	return
}

func (a *Album) ToSearchable() (searchable string) {
	return toSearchable(a.Artist, a.Title, false)
}
