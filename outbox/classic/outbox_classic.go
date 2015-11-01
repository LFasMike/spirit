package classic

import (
	"github.com/gogap/spirit"
	"sync"
	"time"
)

const (
	outboxURN = "urn:spirit:outbox:classic"
)

var _ spirit.Outbox = new(ClassicOutbox)

type ClassicOutboxConfig struct {
	Size       int           `json:"size"`
	GetTimeout int           `json:"get_timeout"`
	Labels     spirit.Labels `json:"labels"`
}

type ClassicOutbox struct {
	statusLocker sync.Mutex
	status       spirit.Status

	senders    []spirit.Sender
	senderLock sync.Mutex

	deliveriesChan chan []spirit.Delivery

	conf ClassicOutboxConfig
}

func init() {
	spirit.RegisterOutbox(outboxURN, NewClassicOutbox)
}

func NewClassicOutbox(options spirit.Options) (box spirit.Outbox, err error) {
	conf := ClassicOutboxConfig{}

	if err = options.ToObject(&conf); err != nil {
		return
	}

	box = &ClassicOutbox{
		conf: conf,
	}
	return

}

func (p *ClassicOutbox) Start() (err error) {
	p.statusLocker.Lock()
	defer p.statusLocker.Unlock()

	if p.status == spirit.StatusRunning {
		return
	}

	spirit.Logger().WithField("actor", "outbox").
		WithField("type", "classic").
		WithField("event", "start").
		Infoln("enter start")

	p.deliveriesChan = make(chan []spirit.Delivery, p.conf.Size)

	p.status = spirit.StatusRunning

	for _, sender := range p.senders {

		go func(sender spirit.Sender) {
			if sender.Status() == spirit.StatusStopped {
				if err = sender.Start(); err != nil {
					spirit.Logger().WithField("actor", "outbox").
						WithField("type", "classic").
						WithField("event", "start sender").
						Errorln(err)
				}

				spirit.Logger().WithField("actor", "outbox").
					WithField("type", "classic").
					WithField("event", "start sender").
					Debugln("sender started")
			}
		}(sender)
	}

	spirit.Logger().WithField("actor", "outbox").
		WithField("type", "classic").
		WithField("event", "start").
		Infoln("started")

	return
}

func (p *ClassicOutbox) Stop() (err error) {
	p.statusLocker.Lock()
	defer p.statusLocker.Unlock()

	if p.status == spirit.StatusStopped {
		return
	}

	spirit.Logger().WithField("actor", "outbox").
		WithField("type", "classic").
		WithField("event", "stop").
		Infoln("enter stop")

	p.status = spirit.StatusStopped

	close(p.deliveriesChan)

	spirit.Logger().WithField("actor", "outbox").
		WithField("type", "classic").
		WithField("event", "stop").
		Infoln("stopped")

	return
}

func (p *ClassicOutbox) Status() (status spirit.Status) {
	return p.status
}

func (p *ClassicOutbox) Labels() (labels spirit.Labels) {
	labels = p.conf.Labels
	return
}

func (p *ClassicOutbox) AddSender(sender spirit.Sender) (err error) {
	p.senderLock.Lock()
	defer p.senderLock.Unlock()

	sender.SetDeliveryGetter(p)
	p.senders = append(p.senders, sender)

	spirit.Logger().WithField("actor", "outbox").
		WithField("type", "classic").
		WithField("event", "add sender").
		Debugln("sender added")

	return
}

func (p *ClassicOutbox) Get() (deliveries []spirit.Delivery, err error) {
	if p.conf.GetTimeout < 0 {
		deliveries = <-p.deliveriesChan
	} else {
		select {
		case deliveries = <-p.deliveriesChan:
			{
				spirit.Logger().WithField("actor", "outbox").
					WithField("type", "classic").
					WithField("event", "get deliveries").
					WithField("length", len(deliveries)).
					Debugln("deliveries received from deliveries chan")
			}
		case <-time.After(time.Duration(p.conf.GetTimeout) * time.Millisecond):
			{
				spirit.Logger().WithField("actor", "outbox").
					WithField("type", "classic").
					WithField("event", "get deliveries").
					Debugln("get deliveries timeout")

				return
			}
		}
	}

	return
}
