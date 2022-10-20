package event

import (
	"encoding/json"
	"fmt"

	"github.com/jellyfishphp/event-queue-worker/internal/process"
)

type Listener struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	EventName  string
}

type ListenerReaderInterface interface {
	GetAll() []Listener
}

type ListenerReader struct {
	command string
}

func (listenerReader *ListenerReader) GetAll() []Listener {
	var ungroupedListeners []Listener

	process := process.NewProcess(listenerReader.command)
	output, error := process.Start()

	if error != nil {
		return ungroupedListeners
	}

	var groupedListeners map[string][]Listener
	error = json.Unmarshal(output, &groupedListeners)

	if error != nil {
		return ungroupedListeners
	}

	for eventName, listeners := range groupedListeners {
		for _, listener := range listeners {
			listener.EventName = eventName
			ungroupedListeners = append(ungroupedListeners, listener)
		}
	}

	return ungroupedListeners
}

func NewListenerReader(pathToConsoleApp string) ListenerReaderInterface {
	return &ListenerReader{
		command: fmt.Sprintf("%s event:listener:get async", pathToConsoleApp),
	}
}

type ListenerConverterInterface interface {
	ConvertToProcess(listener Listener) process.ProcessInterface
	ConvertMultipleToProcesses(listeners []Listener) []process.ProcessInterface
}

type ListenerConverter struct {
	pathToConsoleApp string
}

func (listenerConverter *ListenerConverter) ConvertToProcess(listener Listener) process.ProcessInterface {
	return process.NewProcess(
		fmt.Sprintf(
			"%s event:queue:consume %s %s",
			listenerConverter.pathToConsoleApp,
			listener.EventName,
			listener.Identifier,
		),
	)
}

func (listenerConverter *ListenerConverter) ConvertMultipleToProcesses(listeners []Listener) []process.ProcessInterface {
	var processes []process.ProcessInterface

	for _, listener := range listeners {
		processes = append(processes, listenerConverter.ConvertToProcess(listener))
	}

	return processes
}

func NewListenerConverter(pathToConsoleApp string) ListenerConverterInterface {
	return &ListenerConverter{
		pathToConsoleApp: pathToConsoleApp,
	}
}
