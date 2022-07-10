package ym

import (
	"errors"

	"github.com/oklookat/toanother/core/base"
)

type artistsReponse struct {
	Hooks   *base.Hooks
	Success bool      `json:"success"`
	Artists []*Artist `json:"artists"`
}

// download artists from YM.
func (a *artistsReponse) Download() (response *artistsReponse, err error) {
	if a.Hooks != nil && a.Hooks.OnFetchFromAPI != nil {
		a.Hooks.OnFetchFromAPI(1, 1)
	}

	// prepare request.
	var client = createRequestor(REFERER_PLAYLISTS)
	response = &artistsReponse{}

	// send.
	_, err = client.R().
		SetQueryParam("filter", REFERER_ARTISTS).
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

// track artist.
type Artist struct {
	Hooks *base.Hooks
	Name  string `json:"name"`
}

// download all liked artists.
//
// Then convert it to []base.Playlist & add to DB.
//
// You must make your songs public.
func (a *Artist) DownloadAll() (artists []*base.Artist, err error) {
	var arResp = &artistsReponse{}
	arResp.Hooks = a.Hooks

	response, err := arResp.Download()
	if err != nil {
		return
	}

	// delete artists from DB.
	var tmp = base.Artist{}
	tmp.DeleteAllFromTable(dbConn)

	// convert.
	artists = make([]*base.Artist, 0)
	for i := range response.Artists {
		if a.Hooks != nil && a.Hooks.OnAddingToDatabase != nil {
			a.Hooks.OnAddingToDatabase(i, len(response.Artists))
		}
		baseAr, errd := response.Artists[i].ToBase()
		if errd != nil {
			err = errd
			return
		}
		artists = append(artists, baseAr)
	}

	return
}

// get artists from DB.
func (a *Artist) GetAll() (artists []*base.Artist, err error) {
	if a.Hooks.OnFetchFromDatabase != nil {
		a.Hooks.OnFetchFromDatabase(1, 1)
	}
	var b = base.Artist{}
	return b.GetAll(dbConn)
}

// add to liked artists.
func (a *Artist) ToBase() (artist *base.Artist, err error) {
	artist = &base.Artist{}
	artist.Name = a.Name
	artist.ID, err = artist.AddToTable(dbConn)
	return
}

func (a *Artist) collectNames(ar []*Artist) (names []string, err error) {
	if ar == nil {
		err = errors.New("[artist] empty slice")
		return
	}
	var artists = make([]string, 0)
	for _, pa := range ar {
		artists = append(artists, pa.Name)
	}
	names = artists
	return
}
