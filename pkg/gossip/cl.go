package gossip

import "context"

type Gossipy interface {
	Gossip(context.Context)
}
