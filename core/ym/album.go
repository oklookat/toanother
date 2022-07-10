package ym

import (
	"errors"
	"time"

	"github.com/oklookat/toanother/core/base"
)

type albumsResponse struct {
	Hooks     *base.Hooks
	Success   bool     `json:"success"`
	HasTracks bool     `json:"hasTracks"`
	Albums    []*Album `json:"albums"`
}

// download albums.
func (a *albumsResponse) Download() (response *albumsResponse, err error) {
	if a.Hooks != nil && a.Hooks.OnFetchFromAPI != nil {
		a.Hooks.OnFetchFromAPI(1, 1)
	}

	// prepare request.
	var client = createRequestor(REFERER_ALBUMS)

	response = &albumsResponse{}

	// send.
	_, err = client.R().
		SetQueryParam("filter", REFERER_ALBUMS).
		SetResult(response).
		EnableDump().
		Get(LIBRARY_HANDLER)

	if err != nil {
		return
	}

	if !response.Success {
		err = errors.New("[API] JSON response not success")
		return
	}
	return
}

// album.
type Album struct {
	Hooks *base.Hooks

	// album artists.
	Artists []*Artist `json:"artists"`

	// album ID.
	ID int `json:"id"`

	// example: 2021-01-15T00:00:00+03:00
	ReleaseDate string `json:"releaseDate"`

	// album name.
	Title string `json:"title"`

	// album track count.
	TrackCount int `json:"trackCount"`

	// album year.
	Year int `json:"year"`
}

func (p *Album) DownloadAll() (albums []*base.Album, err error) {
	var alResp = &albumsResponse{
		Hooks: p.Hooks,
	}

	response, err := alResp.Download()
	if err != nil {
		return
	}

	// delete from DB.
	var tmp = base.Album{}
	tmp.DeleteAllFromTable(dbConn)

	// convert.
	albums = make([]*base.Album, 0)
	for _, al := range response.Albums {
		baseAlbum, errd := al.ToBase()
		if errd != nil {
			err = errd
			return
		}
		albums = append(albums, baseAlbum)
	}

	return
}

func (p *Album) GetAll() (albums []*base.Album, err error) {
	if p.Hooks.OnFetchFromDatabase != nil {
		p.Hooks.OnFetchFromDatabase(1, 1)
	}
	var b = base.Album{}
	return b.GetAll(dbConn)
}

// add to liked albums.
func (p *Album) ToBase() (album *base.Album, err error) {
	album = &base.Album{}
	album.Title = p.Title

	parsedTime, err := time.Parse(time.RFC3339, p.ReleaseDate)
	if err != nil {
		err = errors.New("failed to parse releaseDate. Error: " + err.Error())
		return
	}
	album.ReleaseDate = parsedTime.Unix()

	album.TrackCount = p.TrackCount
	album.Year = p.Year

	var ar = Artist{}
	album.Artist, err = ar.collectNames(p.Artists)
	if err != nil {
		return
	}

	album.ID, err = album.AddToTable(dbConn)
	return
}
