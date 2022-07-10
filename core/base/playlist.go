package base

import (
	"database/sql"
)

type Playlist struct {
	ID            int64  `json:"id"`
	IsLikedTracks bool   `json:"isLikedTracks"`
	Title         string `json:"title"`
	TrackCount    int    `json:"trackCount"`
}

// get playlist list.
func (p *Playlist) GetAll(conn *sql.DB) (playlists []*Playlist, err error) {
	rows, err := conn.Query("SELECT * FROM playlist")
	if err != nil {
		return
	}
	defer rows.Close()
	playlists = make([]*Playlist, 0)
	for rows.Next() {
		var p = &Playlist{}
		err = rows.Scan(&p.ID, &p.IsLikedTracks, &p.Title, &p.TrackCount)
		if err != nil {
			return
		}
		playlists = append(playlists, p)
	}
	return
}

// get tracks list by playlist id.
func (p *Playlist) GetTracks(conn *sql.DB) (tracks []*Track, err error) {
	rows, err := conn.Query("SELECT * FROM playlist_track WHERE playlist_id=$1", p.ID)
	if err != nil {
		return
	}
	defer rows.Close()
	tracks = make([]*Track, 0)
	var liked = Artist{}
	for rows.Next() {
		var p = &Track{}
		var artistStr string
		err = rows.Scan(&p.ID, &p.PlaylistID, &p.Title, &p.AlbumTitle, &p.DurationMs, &artistStr)
		if err != nil {
			return
		}
		p.Artist = liked.namesToSlice(artistStr)
		tracks = append(tracks, p)
	}
	return
}

// delete all from table.
func (p *Playlist) DeleteAllFromTable(conn *sql.DB) (err error) {
	_, err = conn.Exec(`DELETE FROM playlist`)
	return
}

// add new to table.
func (p *Playlist) AddToTable(conn *sql.DB) (id int64, err error) {
	result, err := conn.Exec(`INSERT INTO 
	playlist (is_liked_tracks, title, track_count) 
	VALUES ($1, $2, $3)`,
		p.IsLikedTracks, p.Title, p.TrackCount)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}
