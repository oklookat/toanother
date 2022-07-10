package datadir

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGetFullPath(t *testing.T) {
	var res, _ = GetFullPath("./ym/playlists")
	t.Log(res)
}
