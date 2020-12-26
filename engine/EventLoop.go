package engine

import "sync"

type EventLoop struct {
	sync.Mutex
	messageQueue []Command

	postChannel   chan struct{}
	finishChannel chan struct{}

	isReady    bool
	isFinished bool
}

func (eventLoop *EventLoop) getCommand() Command {
	eventLoop.Lock()

	if len(eventLoop.messageQueue) == 0 {
		eventLoop.isReady = true

		eventLoop.Unlock()
		<-eventLoop.postChannel
		eventLoop.Lock()
	}

	command := eventLoop.messageQueue[0]
	eventLoop.messageQueue = eventLoop.messageQueue[1:]
	eventLoop.Unlock()

	return command
}

func (eventLoop *EventLoop) startProcessing() {
	for {
		command := eventLoop.getCommand()
		command.Execute(eventLoop)
		if len(eventLoop.messageQueue) == 0 && eventLoop.isFinished {
			break
		}
	}
	eventLoop.finishChannel <- struct{}{}
}

func (eventLoop *EventLoop) Start() {
	eventLoop.postChannel = make(chan struct{})
	eventLoop.finishChannel = make(chan struct{}, 1)

	go eventLoop.startProcessing()
}

func (eventLoop *EventLoop) Post(command Command) {
	eventLoop.Lock()
	eventLoop.messageQueue = append(eventLoop.messageQueue, command)

	if eventLoop.isReady {
		eventLoop.isReady = false
		eventLoop.postChannel <- struct{}{}
	}
	eventLoop.Unlock()
}

func (eventLoop *EventLoop) AwaitFinish() {
	eventLoop.isFinished = true
	<-eventLoop.finishChannel
}
