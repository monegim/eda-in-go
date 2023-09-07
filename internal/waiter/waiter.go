package waiter

import "context"

type WaitFunc func(context.Context) error
type Waiter interface {
	Add(fns ...WaitFunc)
	Wait() error
	Context() context.Context
	CancelFunc() context.CancelFunc
}
type waiter struct {
	ctx    context.Context
	fns    []WaitFunc
	cancel context.CancelFunc
}
type waiterCfg struct {
	parentCtx    context.Context
	catchSignals bool
}

func New(options ...WaiterOption) Waiter {

}
