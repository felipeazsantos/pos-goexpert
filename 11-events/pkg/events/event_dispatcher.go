package events

import (
	"errors"
	"slices"
)

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

func (d *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if handlers, exists := d.handlers[eventName]; exists {
		for i, h := range handlers {
			if h == handler {
				d.handlers[eventName] = append(handlers[:i], handlers[i+1:]...)
				return nil
			}
		}
		return errors.New("handler not found")
	}
	return errors.New("event not found")
}

func (d *EventDispatcher) Clear() {
	for eventName := range d.handlers {
		delete(d.handlers, eventName)
	}
}

func (d *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if handlers, exists := d.handlers[eventName]; exists {
		return slices.Contains(handlers, handler)
	}
	return false
}

func (d *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, exists := d.handlers[event.GetName()]; exists {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	return nil
}