package base

import (
	"database/sql"
	"strings"
)

type Artist struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// get artist list.
func (l *Artist) GetAll(conn *sql.DB) (artists []*Artist, err error) {
	rows, err := conn.Query("SELECT * FROM liked_artist")
	if err != nil {
		return
	}
	defer rows.Close()
	artists = make([]*Artist, 0)
	for rows.Next() {
		var a = &Artist{}
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return
		}
		artists = append(artists, a)
	}
	return
}

// delete all from table.
func (l *Artist) DeleteAllFromTable(conn *sql.DB) (err error) {
	_, err = conn.Exec(`DELETE FROM liked_artist`)
	return
}

// add new to table.
func (a *Artist) AddToTable(conn *sql.DB) (id int64, err error) {
	result, err := conn.Exec(`INSERT INTO 
	liked_artist (name) 
	VALUES ($1)`,
		a.Name)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}

func (a *Artist) namesToSlice(names string) []string {
	return strings.Split(names, ",")
}

func (a *Artist) sliceToNames(names []string) string {
	if names == nil || len(names) < 1 {
		return ""
	}
	return strings.Join(names, ", ")
}
