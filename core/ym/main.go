package ym

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/imroc/req/v3"
	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/database"
	"github.com/oklookat/toanother/core/datadir"
)

const (
	// other.
	USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"

	// dirs.
	WORKDIR   = "."
	DUMP_NAME = WORKDIR + "/ym_dump.txt"
	DB_NAME   = WORKDIR + "/ym.db"

	// playlist.
	PLAYLIST_HANDLER  = "https://music.yandex.ru/handlers/playlist.jsx"
	REFERER_PLAYLISTS = "playlists"

	// library.
	LIBRARY_HANDLER = "https://music.yandex.ru/handlers/library.jsx"
	REFERER_ARTISTS = "artists"
	REFERER_ALBUMS  = "albums"
)

var Settings *base.YandexMusicSettings
var dbConn *sql.DB

func Init() (err error) {
	Settings = &base.ConfigFile.YandexMusic
	dbConn, err = database.Load(DB_NAME, func(db *sql.DB) error {
		dbConn = db
		return base.RecreateAll(dbConn)
	})
	if err != nil {
		return
	}
	return
}

func createRequestor(referer string) *req.Client {
	referer = fmt.Sprintf("https://music.yandex.ru/users/%v/"+referer, Settings.Login)
	var client = req.C().
		SetUserAgent(USER_AGENT).
		SetTimeout(15 * time.Second).
		SetJsonUnmarshal(func(data []byte, v interface{}) error {
			// write response json to file.
			var tempFile, errd = datadir.TempFile()
			if errd != nil {
				return errd
			}

			defer func() {
				_ = tempFile.Close()
				_ = os.Remove(tempFile.Name())
			}()

			if _, errd = tempFile.Write(data); errd != nil {
				return errd
			}

			if _, errd = tempFile.Seek(0, io.SeekStart); errd != nil {
				return errd
			}

			// convert json to struct.
			return json.NewDecoder(tempFile).Decode(v)
		})
	//.DevMode()
	client.SetCommonHeaders(map[string]string{
		"Referer":          referer,
		"X-Retpath-Y":      referer,
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Language":  "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7",
		"Connection":       "keep-alive",
		"X-Requested-With": "XMLHttpRequest",
	})
	client.SetCommonQueryParams(map[string]string{
		"external-domain": "music.yandex.ru",
		"overembed":       "false",
		"lang":            "ru",
		"owner":           Settings.Login,
		"likeFilter":      "favorite",
	})

	client.OnAfterResponse(func(c *req.Client, r *req.Response) (err error) {
		if !r.IsSuccess() {
			err = errors.New(r.Dump())
			return
		}
		return
	})

	return client
}
