package events

import "errors"

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}
func (d *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, exists := d.handlers[eventName]; exists {
		for _, h := range d.handlers[eventName] {
			if h == handler {
				return errors.New("handler already registered")
			}
		}
	}

	d.handlers[eventName] = append(d.handlers[eventName], handler)
	return nil
}