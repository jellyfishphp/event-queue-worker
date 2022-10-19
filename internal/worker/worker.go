package worker

import (
	"github.com/jellyfishphp/event-queue-worker/internal/process"
)

type WorkerInterface interface {
	Start() error
}

type EventQueueWorker struct {
	processList []process.ProcessInterface
}

func NewEventQueueWorker(processes []process.ProcessInterface) WorkerInterface {
	return &EventQueueWorker{
		processList: processes,
	}
}

func (worker *EventQueueWorker) Start() error {
	for {
		for _, currentProcess := range worker.processList {
			go func(process process.ProcessInterface) {
				if process.IsRunning() {
					return
				}

				process.Start()
			}(currentProcess)
		}
	}
}
