package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jellyfishphp/event-queue-worker/internal/event"
	"github.com/jellyfishphp/event-queue-worker/internal/worker"
)

func main() {
	jellyfishApplicationDir, error := os.Getwd()

	if error != nil {
		log.Fatal(error)
	}

	if len(os.Args) == 2 {
		jellyfishApplicationDir = os.Args[1]
	}

	consoleCommand := fmt.Sprintf("%s/vendor/bin/console", strings.TrimRight(jellyfishApplicationDir, "/"))

	listenerReader := event.NewListenerReader(consoleCommand)
	listeners := listenerReader.GetAll()
	listenerConverter := event.NewListenerConverter(consoleCommand)
	processes := listenerConverter.ConvertMultipleToProcesses(listeners)
	eventQueueWorker := worker.NewEventQueueWorker(processes)

	eventQueueWorker.Start()
}
