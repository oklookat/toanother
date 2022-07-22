package spotify

import (
	"context"

	"github.com/zmb3/spotify/v2"
)

type searcher interface {
	Search(ctx context.Context, query string, t spotify.SearchType, opts ...spotify.RequestOption) (*spotify.SearchResult, error)
}

type trackAdder interface {
	AddTracksToLibrary(ctx context.Context, ids ...spotify.ID) error
}

type albumAdder interface {
	AddAlbumsToLibrary(ctx context.Context, ids ...spotify.ID) error
}

type artistFollower interface {
	FollowArtist(ctx context.Context, ids ...spotify.ID) error
}

type currentUsersAlbums interface {
	CurrentUsersAlbums(ctx context.Context, opts ...spotify.RequestOption) (*spotify.SavedAlbumPage, error)
}

type albumsRemover interface {
	RemoveAlbumsFromLibrary(ctx context.Context, ids ...spotify.ID) error
}
