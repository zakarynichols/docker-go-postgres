package app

import "context"

type Pinger interface {
	Ping(context.Context) (string, error)
}
