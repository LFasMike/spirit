package spirit

import (
	"sync"
)

var (
	outboxesLocker = sync.Mutex{}
	newOutboxFuncs = make(map[string]NewOutboxFunc)
)

type NewOutboxFunc func(config Config) (outbox Outbox, err error)

type Outbox interface {
	StartStopper

	Labels() Labels

	DeliveryGetter
	DeliveryPutter
}

func RegisterOutbox(urn string, newFunc NewOutboxFunc) (err error) {
	outboxesLocker.Lock()
	outboxesLocker.Unlock()

	if urn == "" {
		panic("spirit: Register outbox urn is empty")
	}

	if newFunc == nil {
		panic("spirit: Register outbox is nil")
	}

	if _, exist := newOutboxFuncs[urn]; exist {
		panic("spirit: Register called twice for outbox " + urn)
	}

	newOutboxFuncs[urn] = newFunc

	logger.WithField("module", "spirit").
		WithField("event", "register outbox").
		WithField("urn", urn).
		Debugln("outbox registered")

	return
}
