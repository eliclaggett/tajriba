package runtime

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/sasha-s/go-deadlock"
)

type Runtime struct {
	*objectMap
	stepTimers  map[string]*time.Timer
	changesSubs map[string][]*changesSub
	onEventSubs map[string][]*onEventSub
	sattrSubs   map[string][]*scopedAttributesSub

	deadlock.RWMutex
}

func Start(ctx context.Context) (*Runtime, error) {
	log.Debug().Msg("runtime: started")

	r := &Runtime{
		objectMap:   newObjectMap(),
		stepTimers:  make(map[string]*time.Timer),
		changesSubs: make(map[string][]*changesSub),
		onEventSubs: make(map[string][]*onEventSub),
		sattrSubs:   make(map[string][]*scopedAttributesSub),
	}

	if err := r.load(ctx); err != nil {
		return nil, errors.Wrap(err, "load values")
	}

	return r, nil
}

func (r *Runtime) Stop() {
}