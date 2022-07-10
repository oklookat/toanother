package ym

import (
	"errors"
	"strconv"
	"time"

	"github.com/oklookat/toanother/core/base"
)

// playlists list with small information.
type playlistsResponse struct {
	Hooks       *base.Hooks
	Success     bool  `json:"success"`
	HasTracks   bool  `json:"hasTracks"`
	PlaylistIds []int `json:"playlistIds"`
}

// download playlists ID's.
func (p *playlistsResponse) Download() (response *playlistsResponse, err error) {
	if p.Hooks != nil && p.Hooks.OnFetchFromAPI != nil {
		p.Hooks.OnFetchFromAPI(1, 1)
	}

	// prepare request.
	var client = createRequestor(REFERER_PLAYLISTS)

	response = &playlistsResponse{}

	// send.
	_, err = client.R().
		SetQueryParam("filter", REFERER_PLAYLISTS).
		SetQueryParam("playlistsWithoutContent", "true").
		SetQueryParam("likedPlaylistsPage", "0").
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

type Playlist struct {
	Hooks *base.Hooks
	// id of playlist.
	Kind       int      `json:"kind"`
	Title      string   `json:"title"`
	TrackCount int      `json:"trackCount"`
	Tracks     []*Track `json:"tracks"`
}

// download all playlists.
//
// Then convert it to []base.Playlist & add to DB.
//
// You must make your songs public.
func (p *Playlist) DownloadAll() (playlists []*base.Playlist, err error) {
	// download playlists id's.
	var plResp = &playlistsResponse{}
	plResp.Hooks = p.Hooks

	response, err := plResp.Download()
	if err != nil {
		return
	}

	// delete playlists from DB.
	var tmp = base.Playlist{}
	tmp.DeleteAllFromTable(dbConn)

	// convert playlists id's to playlist.
	playlists = make([]*base.Playlist, 0)
	for i, id := range response.PlaylistIds {
		if p.Hooks != nil && p.Hooks.OnFetchFromAPI != nil {
			p.Hooks.OnFetchFromAPI(i, len(response.PlaylistIds))
		}
		basePl, errd := p.Download(strconv.Itoa(id))
		if errd != nil {
			err = errd
			return
		}
		playlists = append(playlists, basePl)
		if i < len(response.PlaylistIds) {
			time.Sleep(2 * time.Second)
		}
	}
	return
}

// download playlist by YM playlist id.
//
// Then convert it to base.Playlist & add to DB.
//
// You must make your songs public.
func (p *Playlist) Download(playlistID string) (playlist *base.Playlist, err error) {
	var playlistResponse *struct {
		Playlist Playlist `json:"playlist"`
	} = &struct {
		Playlist Playlist "json:\"playlist\""
	}{}

	// prepare request.
	var client = createRequestor(REFERER_PLAYLISTS)

	// send request.
	_, err = client.R().
		SetQueryParam("owner", Settings.Login).SetQueryParam("kinds", playlistID).
		SetResult(playlistResponse).
		EnableDump().
		Get(PLAYLIST_HANDLER)

	// check.
	if err != nil {
		return
	}

	// convert.
	playlistResponse.Playlist.Hooks = p.Hooks
	return playlistResponse.Playlist.toBase()
}

// get playlists from DB.
func (p *Playlist) GetPlaylists() (playlists []*base.Playlist, err error) {
	if p.Hooks.OnFetchFromDatabase != nil {
		p.Hooks.OnFetchFromDatabase(1, 1)
	}
	var b = base.Playlist{}
	return b.GetAll(dbConn)
}

// get playlist tracks from DB.
func (p *Playlist) GetTracks(playlistID int64) (tracks []*base.Track, err error) {
	if p.Hooks.OnFetchFromDatabase != nil {
		p.Hooks.OnFetchFromDatabase(1, 1)
	}
	var b = base.Playlist{}
	b.ID = playlistID
	return b.GetTracks(dbConn)
}

// convert YM playlist to base.Playlist.
func (p *Playlist) toBase() (basePlaylist *base.Playlist, err error) {
	basePlaylist = &base.Playlist{}
	basePlaylist.IsLikedTracks = p.Kind == 3
	basePlaylist.Title = p.Title
	basePlaylist.TrackCount = p.TrackCount
	basePlaylist.ID, err = basePlaylist.AddToTable(dbConn)
	if err != nil {
		return
	}

	if p.Tracks != nil {
		for i, track := range p.Tracks {
			if p.Hooks != nil && p.Hooks.OnAddingToDatabase != nil {
				p.Hooks.OnAddingToDatabase(i, len(p.Tracks))
			}
			track.ToBase(basePlaylist.ID)
		}
	}

	return
}
