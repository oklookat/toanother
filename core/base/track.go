package base

import (
	"database/sql"
	"errors"
	"fmt"
)

type Track struct {
	ID         int64    `json:"id"`
	PlaylistID int64    `json:"playlistID"`
	Title      string   `json:"title"`
	AlbumTitle *string  `json:"albumTitle"`
	DurationMs int      `json:"durationMs"`
	Artist     []string `json:"artist"`
}

// add new to table.
func (t *Track) AddToTable(conn *sql.DB) (id int64, err error) {
	if t.Artist == nil {
		err = errors.New("empty artist")
		return
	}
	var liked = Artist{}
	var artistConv = liked.sliceToNames(t.Artist)

	result, err := conn.Exec(`INSERT INTO 
	playlist_track (playlist_id, title, album_title, duration_ms, artist) 
	VALUES ($1, $2, $3, $4, $5)`,
		t.PlaylistID, t.Title, t.AlbumTitle, t.DurationMs, artistConv)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}

// convert track title to searchable string.
func (t *Track) ToSearchable() (searchable string) {
	if len(t.Artist) < 1 {
		return
	}

	// modify artist.
	var artistStr = t.Artist[0]
	artistStr = REGEXP_SYMBOLS.ReplaceAllString(artistStr, "")

	// modify title.
	var title = t.Title
	title = REGEXP_BRACKETS.ReplaceAllString(title, "")
	title = REGEXP_SYMBOLS.ReplaceAllString(title, "")

	searchable = fmt.Sprintf("%v %v", artistStr, title)
	return
}
