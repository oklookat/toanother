package spotify

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/oklookat/toanother/core/datadir"
	"github.com/oklookat/toanother/core/logger"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

// auth.
func (i *Instance) WebAuth() (err error) {
	var httpErr = make(chan error, 1)
	if !i.isWebAuthCalledBefore {
		http.HandleFunc(spotifyURI, func(w http.ResponseWriter, r *http.Request) {
			var errd error
			i.token, errd = i.authenticator.Token(r.Context(), i.state, r)

			if errd != nil {
				http.Error(w, "couldn't get token", http.StatusForbidden)
				httpErr <- errd
				return
			}

			if st := r.FormValue("state"); st != i.state {
				var errStr = fmt.Sprintf("state mismatch: %s != %s\n", st, i.state)
				http.Error(w, errStr, 400)
				httpErr <- errors.New(errStr)
				return
			}

			// use the token to get an authenticated client
			errd = i.authByToken(r.Context())

			httpErr <- errd
		})

		go func() {
			if err = http.ListenAndServe(":8080", nil); err != nil {
				logger.Log.Fatal(err.Error())
			}
		}()

		i.isWebAuthCalledBefore = true
	}

	var url = i.authenticator.AuthURL(i.state)
	if i.hooks != nil && i.hooks.OnAuthURL != nil {
		i.hooks.OnAuthURL(url)
	}

	// wait error handler (auth complete).
	if err = <-httpErr; err != nil {
		return err
	}

	if err = i.WriteToken(i.token); err != nil {
		return
	}

	return
}

func (i *Instance) authByToken(ctx context.Context) (err error) {
	i.client = spotify.New(i.authenticator.Client(ctx, i.token))
	// use the client to make calls that require authorization
	i.user, err = i.client.CurrentUser(context.Background())
	return
}

// read token from json.
func (i *Instance) readToken() (err error) {
	isExists, err := datadir.IsFileExists(TOKEN_DIR)
	if err != nil {
		return
	}
	if !isExists {
		return
	}
	i.token = &oauth2.Token{}
	err = datadir.GetStructByFile(TOKEN_DIR, false, i.token)
	return
}

// write token to json.
func (i *Instance) WriteToken(token *oauth2.Token) (err error) {
	if token == nil {
		err = errors.New("nil token")
		return
	}
	return datadir.WriteFileStruct(TOKEN_DIR, false, token)
}

func (i *Instance) Ping() (err error) {
	if i.client == nil || i.token == nil {
		return
	}
	usr, err := i.client.CurrentUser(context.Background())
	if err != nil {
		return
	}
	if len(usr.ID) < 1 {
		err = errors.New("spotify: ping failed")
	}
	return
}
