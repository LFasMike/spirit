package spirit

import (
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/gogap/ali_mqs"
	"github.com/gogap/errors"
	"github.com/gogap/logs"
)

type MessageReceiverMQS struct {
	url string

	queue ali_mqs.AliMQSQueue

	recvLocker  sync.Mutex
	isReceiving bool

	responseChan chan ali_mqs.MessageReceiveResponse
	errorChan    chan error

	isPaused bool
}

func NewMessageReceiverMQS(url string) MessageReceiver {
	return &MessageReceiverMQS{url: url}
}

func (p *MessageReceiverMQS) Init(url, configFile string) (err error) {
	p.url = url

	var queue ali_mqs.AliMQSQueue
	if queue, err = p.newAliMQSQueue(); err != nil {
		return
	}

	p.queue = queue

	return
}

func (p *MessageReceiverMQS) Type() string {
	return "mqs"
}

func (p *MessageReceiverMQS) Address() MessageAddress {
	return MessageAddress{Type: p.Type(), Url: p.url}
}

func (p *MessageReceiverMQS) newAliMQSQueue() (queue ali_mqs.AliMQSQueue, err error) {

	hostId := ""
	accessKeyId := ""
	accessKeySecret := ""
	queueName := ""

	regUrl := regexp.MustCompile("(.*):(.*)@(.*)/(.*)")
	regMatched := regUrl.FindAllStringSubmatch(p.url, -1)

	if len(regMatched) == 1 &&
		len(regMatched[0]) == 5 {
		accessKeyId = regMatched[0][1]
		accessKeySecret = regMatched[0][2]
		hostId = regMatched[0][3]
		queueName = regMatched[0][4]
	}

	client := ali_mqs.NewAliMQSClient(hostId,
		accessKeyId,
		accessKeySecret)

	if client == nil {
		err = ERR_RECEIVER_MQS_CLIENT_IS_NIL.New(errors.Params{"type": p.Type(), "url": p.url})
		return
	}

	queue = ali_mqs.NewMQSQueue(queueName, client)

	return
}

func (p *MessageReceiverMQS) Receive(portChan *PortChan) {
	p.recvLocker.Lock()
	defer p.recvLocker.Unlock()
	if p.isReceiving == true {
		panic("could not start receive twice")
	}
	p.isReceiving = true

	recvLoop := func(receiverType string, url string, queue ali_mqs.AliMQSQueue, compMsgChan chan ComponentMessage, compErrChan chan error, singalChan chan int, stoppedChan chan bool) {
		responseChan := make(chan ali_mqs.MessageReceiveResponse, MESSAGE_CHAN_SIZE)
		errorChan := make(chan error, MESSAGE_CHAN_SIZE)

		defer close(responseChan)
		defer close(errorChan)

		go queue.ReceiveMessage(responseChan, errorChan)
		logs.Info("receiver listening - ", p.url)

		isPaused := false
		isStopping := false

		stoplogTime := time.Now()
		for {
			if !isStopping {
				select {
				case signal := <-singalChan:
					{
						switch signal {
						case SIG_PAUSE:
							stopQueue(queue)
							logs.Warn(fmt.Sprintf("* mqs receiver of %s paused - resp chan len: %d, err chan len: %d", queue.Name(), len(responseChan), len(errorChan)))
							isPaused = true
						case SIG_RESUME:
							logs.Warn("* mqs receiver resumed")
							go queue.ReceiveMessage(responseChan, errorChan)
							isPaused = false
						case SIG_STOP:
							stopQueue(queue)
							logs.Warn(fmt.Sprintf("* mqs receiver of %s stopping - resp chan len: %d, err chan len: %d", queue.Name(), len(responseChan), len(errorChan)))
							isStopping = true
							stoplogTime = time.Now()
						}
					}
				default:
				}
			}

			if isStopping {
				if len(responseChan) == 0 && len(errorChan) == 0 {
					logs.Warn(fmt.Sprintf("* mqs receiver of %s stopped - resp chan len: %d, err chan len: %d", queue.Name(), 0, 0))
					stoppedChan <- true
					return
				} else {
					now := time.Now()
					if now.Sub(stoplogTime) >= 1*time.Millisecond {
						stoplogTime = now
						logs.Warn(fmt.Sprintf("* mqs receiver of %s stopping - resp chan len: %d, err chan len: %d", queue.Name(), len(responseChan), len(errorChan)))
					}
				}
			} else if isPaused {
				if len(responseChan) > 0 || len(errorChan) > 0 {
					logs.Info("[mqs receiver pausing] resp chan len:", len(responseChan))
				}
				time.Sleep(time.Second)
				continue
			}

			select {
			case resp := <-responseChan:
				{
					go func(receiverType string,
						url string,
						queue ali_mqs.AliMQSQueue,
						compMsgChan chan ComponentMessage,
						compErrChan chan error,
						resp ali_mqs.MessageReceiveResponse) {

						if resp.MessageBody != nil && len(resp.MessageBody) > 0 {
							compMsg := ComponentMessage{}
							if e := compMsg.UnSerialize(resp.MessageBody); e != nil {
								e = ERR_RECEIVER_UNMARSHAL_MSG_FAILED.New(errors.Params{"type": receiverType, "url": url, "err": e})
								compErrChan <- e
							}
							compMsgChan <- compMsg
						}

						if e := queue.DeleteMessage(resp.ReceiptHandle); e != nil {
							e = ERR_RECEIVER_DELETE_MSG_ERROR.New(errors.Params{"type": receiverType, "url": url, "err": e})
							compErrChan <- e
						}
					}(receiverType, url, queue, compMsgChan, compErrChan, resp)
				}
			case respErr := <-errorChan:
				{
					go func(err error) {
						if !ali_mqs.ERR_MQS_MESSAGE_NOT_EXIST.IsEqual(err) {
							compErrChan <- err
						}
					}(respErr)
				}
			case <-time.After(time.Second):
				{
					if len(responseChan) == 0 && len(errorChan) == 0 && isStopping {
						stoppedChan <- true
						return
					}
				}
			}

		}
	}
	recvLoop(p.Type(), p.url, p.queue, portChan.Message, portChan.Error, portChan.Signal, portChan.Stoped)
}

func stopQueue(queue ali_mqs.AliMQSQueue) {
	logs.Warn(fmt.Sprintf("* mqs receiver of %s stopping - begin stop receive message from ali-mqs queue", queue.Name()))
	queue.Stop()
	logs.Warn(fmt.Sprintf("* mqs receiver of %s stopping - stoped receive messages from ali-mqs queue", queue.Name()))
}
