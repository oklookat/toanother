package spotify

import (
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/utils"
	"github.com/zmb3/spotify/v2"
)

// ID finding methods.
type idFinder[T base.Entities] interface {
	// find id by this method.
	Find(val T) (found bool, id spotify.ID, err error)
}

func findIds[T base.Entities](vals []T, finder idFinder[T], onProcessing func(current int, total int)) (sIds [][]spotify.ID, err error) {
	if vals == nil {
		err = errors.New("nil vals")
		return
	}
	if finder == nil {
		err = errors.New("nil finder")
		return
	}

	var ids = make([]spotify.ID, 0)
	var total = len(vals)

	for counter, val := range vals {
		if val == nil {
			continue
		}
		if onProcessing != nil {
			onProcessing(counter, total)
		}

		found, foundID, errd := finder.Find(val)
		if errd != nil {
			err = errd
			return
		}
		if !found {
			continue
		}

		ids = append(ids, foundID)
	}

	// split because spotify have limit API calls.
	sIds = utils.SplitSlice(ids, 20)
	return
}
