package spotify

import (
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/zmb3/spotify/v2"
)

// ID finding methods.
type idFinder[T comparable] interface {
	// find ID.
	Find(val T) (found bool, id spotify.ID, err error)

	// (base.Hooks).
	OnImport(current int, total int, notFound []any)

	// when ID's found.
	OnFound(ids [][]spotify.ID) (err error)
}

type findIdsArgs[T comparable] struct {
	instance *Instance
	vals     []T
	finder   idFinder[T]
	hooks    *base.Hooks
}

// Find Spotify ID's.
func findIds[T comparable](args *findIdsArgs[T]) (err error) {
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
	var notFound = 0
	var notFounds = make([]any, 0)
	var total = len(args.vals)

	for counter, val := range args.vals {
		args.finder.OnImport(counter, total, notFounds)
		found, foundID, errd := args.finder.Find(val)
		if errd != nil {
			err = errd
			return
		}
		if !found {
			notFounds = append(notFounds, val)
			notFound++
			continue
		}
		// limit: 50 id's per API call.
		if counter%20 == 0 && len(ids[delimIndex]) > 19 {
			ids = append(ids, make([]spotify.ID, 0))
			delimIndex++
		}
		ids[delimIndex] = append(ids[delimIndex], foundID)
	}

	return args.finder.OnFound(ids)
}
