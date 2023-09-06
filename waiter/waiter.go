package waiter

import "context"

type WaitFunc func(context.Context) error
type Waiter interface {
	Add(fns ...WaitFunc)
	Wait() error
	Context() context.Context
	CancelFunc() context.CancelFunc
}
