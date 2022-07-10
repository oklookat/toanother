package base

import (
	"database/sql"
	"regexp"
)

var (
	// not-searchable symbols.
	REGEXP_SYMBOLS = regexp.MustCompile("[.]+")
	// brackets ( and values inside ).
	REGEXP_BRACKETS = regexp.MustCompile(`(?s)\((.*)\)`)
)

const (
	SQL_PLAYLIST_DROP = `
	DROP TABLE IF EXISTS playlist;
	DROP TABLE IF EXISTS playlist_track;
	DROP TABLE IF EXISTS liked_album;
	DROP TABLE IF EXISTS liked_artist;
	`
	SQL_PLAYLIST = `
	CREATE TABLE playlist(
		id INTEGER PRIMARY KEY, 
		is_liked_tracks INTEGER,
		title TEXT NOT NULL,
		track_count INTEGER
	);
	CREATE TABLE playlist_track(
		id INTEGER PRIMARY KEY,
		playlist_id INTEGER REFERENCES playlist(id) ON DELETE CASCADE NOT NULL,
		title TEXT NOT NULL,
		album_title TEXT,
		duration_ms INTEGER NOT NULL,
		artist TEXT NOT NULL
	);

	CREATE TABLE liked_album(
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		artist TEXT NOT NULL,
		release_date INTEGER NOT NULL,
		track_count INTEGER NOT NULL,
		year INTEGER NOT NULL
	);

	CREATE TABLE liked_artist(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);
	`
)

type Hooks struct {
	// when API request.
	OnFetchFromAPI func(current int, total int)
	// when DB request.
	OnFetchFromDatabase func(current int, total int)
	// when adding to DB.
	OnAddingToDatabase func(current int, total int)
	// on importing tracks in service.
	OnImport func(current int, total int, notFound []interface{})
}

type NotFounder interface {
	NotFoundMsg() string
}

// drop all tables, and create.
func RecreateAll(conn *sql.DB) (err error) {
	if _, err = conn.Exec(SQL_PLAYLIST_DROP); err != nil {
		return err
	}
	_, err = conn.Exec(SQL_PLAYLIST)
	return
}
