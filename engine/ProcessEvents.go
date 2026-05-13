package engine

import (
	"ltz/keys"
	"ltz/shared"
)

func ProcessEvents(events chan shared.Event) {
	for event := range events {
		if event.Type == shared.ENUM_EVENT_KEY {
			if event.KeyData.Key.Equals(keys.And(keys.CTRL, keys.C)) {
				return
			}
		}
	}
}