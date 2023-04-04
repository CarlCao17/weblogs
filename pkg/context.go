package pkg

import (
	"context"
	"time"
)

type RPCContext struct {
	context.Context
	Start     time.Time
	MileStone time.Time
	Method    string
	Extra     map[string]any
}
