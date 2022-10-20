package event

import (
	"fmt"
	"testing"
)

func TestGetAll(t *testing.T) {
	listenerReader := NewListenerReader("./console.sh")
	listeners := listenerReader.GetAll()

	if len(listeners) != 2 {
		t.Errorf("Number of listeners is incorect. Acutal: %d, Expected: %d", len(listeners), 2)
	}

	for index, listener := range listeners {
		expectedEventName := fmt.Sprintf("event%d", index+1)
		expectedIdentifier := fmt.Sprintf("identifier%d", index+1)
		expectedType := "async"

		if listener.EventName != expectedEventName {
			t.Errorf("Event name of listener is incorect. Acutal: %s, Expected: %s", listener.EventName, expectedEventName)
		}

		if listener.Identifier != expectedIdentifier {
			t.Errorf("Identifier of listener is incorect. Acutal: %s, Expected: %s", listener.Identifier, expectedIdentifier)
		}

		if listener.Type != expectedType {
			t.Errorf("Type of the first listener is incorect. Acutal: %s, Expected: %s", listener.Type, expectedType)
		}
	}

}
