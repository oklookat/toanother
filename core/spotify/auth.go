package spotify

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/datadir"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

const (
	AUTH_STATE         = "1234"
	AUTH_CALLBACK_PATH = "spotify/callback"
	AUTH_REDIRECT_URL  = "http://localhost:8080/" + AUTH_CALLBACK_PATH
)

type authHandler struct {
	initialized   bool
	authenticator *spotifyauth.Authenticator
	HttpErr       chan error
	Token         *oauth2.Token
}

func (a *authHandler) New(au *spotifyauth.Authenticator) {
	a.authenticator = au
	a.HttpErr = make(chan error)
	a.initialized = true
}

func (a *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !a.initialized {
		http.Error(w, "authHandler not initialized", 500)
		return
	}

	if r.URL.Path != AUTH_CALLBACK_PATH {
		w.WriteHeader(200)
		return
	}

	var token, err = a.authenticator.Token(r.Context(), AUTH_STATE, r)

	if err != nil {
		http.Error(w, "couldn't get token", http.StatusForbidden)
		a.HttpErr <- err
		return
	}

	if st := r.FormValue("state"); st != AUTH_STATE {
		var errStr = fmt.Sprintf("state mismatch: %s != %s\n", st, AUTH_STATE)
		http.Error(w, errStr, 400)
		a.HttpErr <- errors.New(errStr)
		return
	}

	a.Token = token
	a.HttpErr <- err
}

type auth struct {
	Client        *spotify.Client
	User          *spotify.PrivateUser
	token         *oauth2.Token
	authenticator *spotifyauth.Authenticator
}

// Check auth.
func (a *auth) Ping() (err error) {
	if a.Client == nil || a.token == nil {
		err = errors.New("not authorized")
		return
	}
	usr, err := a.Client.CurrentUser(context.Background())
	if err != nil {
		return
	}
	if len(usr.ID) < 1 {
		err = errors.New("empty user ID")
	}
	return
}

// Get token by web & make auth.
//
// onURL: when we get URL to auth in app.
func (a *auth) Web(onURL func(url string)) (err error) {
	if a.authenticator == nil {
		a.initAuthenticator()
	}

	var handler = &authHandler{}
	handler.New(a.authenticator)

	var server = http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	defer server.Close()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	defer listener.Close()

	go func() {
		_ = server.Serve(listener)
	}()

	var url = a.authenticator.AuthURL(AUTH_STATE)
	if onURL != nil {
		go onURL(url)
	}

	// wait error handler (auth complete).
	if err = <-handler.HttpErr; err != nil {
		return
	}

	// set.
	if err = a.setToken(handler.Token); err != nil {
		return
	}

	// auth.
	if err = a.ByToken(context.Background()); err != nil {
		return
	}

	return
}

// Create client by a.token.
func (a *auth) ByToken(ctx context.Context) (err error) {
	if a.token == nil {
		if ok, errd := a.readTokenFromFile(); !ok || errd != nil {
			return
		}
	}
	a.initAuthenticator()
	a.Client = spotify.New(a.authenticator.Client(ctx, a.token))
	a.User, err = a.Client.CurrentUser(context.Background())
	return
}

// Set a.token from file.
func (a *auth) readTokenFromFile() (ok bool, err error) {
	isExists, err := datadir.IsFileExists(TOKEN_DIR)
	if err != nil {
		return
	}
	if !isExists {
		return
	}
	a.token = &oauth2.Token{}
	if err = datadir.GetStructByFile(TOKEN_DIR, false, a.token); err != nil {
		return
	}
	ok = true
	a.initAuthenticator()
	return
}

// Set a.token & set to file.
func (a *auth) setToken(token *oauth2.Token) (err error) {
	if token == nil {
		err = errors.New("nil token")
		return
	}
	a.token = token
	a.initAuthenticator()
	return datadir.WriteFileStruct(TOKEN_DIR, false, a.token)
}

// Re(init) app authenticator.
func (a *auth) initAuthenticator() {
	a.authenticator = spotifyauth.New(
		spotifyauth.WithClientID(base.ConfigFile.Spotify.ID),
		spotifyauth.WithClientSecret(base.ConfigFile.Spotify.Secret),
		spotifyauth.WithRedirectURL(AUTH_REDIRECT_URL),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserLibraryRead,
			spotifyauth.ScopeUserLibraryModify,
			spotifyauth.ScopeUserFollowRead,
			spotifyauth.ScopeUserFollowModify,
			spotifyauth.ScopePlaylistReadPrivate,
			spotifyauth.ScopePlaylistModifyPrivate,
			spotifyauth.ScopeUserReadPrivate,
		),
	)
}
