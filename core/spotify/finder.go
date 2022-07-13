package spotify

import (
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/zmb3/spotify/v2"
)

type idFinderAllowed interface {
	*base.Track | *base.Artist | *base.Album
}

// ID finding methods.
type idFinder[T idFinderAllowed] interface {
	Find(val T) (found bool, id spotify.ID, err error)

	OnFinish(ids [][]spotify.ID) (err error)
}

type findIdsArgs[T idFinderAllowed] struct {
	instance *Instance
	vals     []T
	finder   idFinder[T]
	hooks    *base.Hooks[T]
}

// Find Spotify ID's by args.
func findIds[T idFinderAllowed](args *findIdsArgs[T]) (err error) {
	if args == nil {
		err = errors.New("nil args")
		return
	}
	if args.instance == nil {
		err = errors.New("nil instance")
		return
	}
	if err = args.instance.Ping(); err != nil {
		return
	}
	if args.finder == nil {
		err = errors.New("nil finder")
		return
	}
	if args.vals == nil {
		err = errors.New("nil vals")
		return
	}

	var ids = make([][]spotify.ID, 0)
	ids = append(ids, make([]spotify.ID, 0))
	var delimIndex = 0
	//
	var total = len(args.vals)

	for counter, val := range args.vals {
		if args.hooks != nil && args.hooks.OnProcessing != nil {
			args.hooks.OnProcessing(counter, total)
		}
		found, foundID, errd := args.finder.Find(val)
		if errd != nil {
			err = errd
			return
		}
		if !found {
			if args.hooks != nil && args.hooks.OnNotFound != nil {
				args.hooks.OnNotFound(val)
			}
			continue
		}
		// limit: 50 id's per API call.
		if counter%20 == 0 && len(ids[delimIndex]) > 19 {
			ids = append(ids, make([]spotify.ID, 0))
			delimIndex++
		}
		ids[delimIndex] = append(ids[delimIndex], foundID)
	}

	return args.finder.OnFinish(ids)
}
