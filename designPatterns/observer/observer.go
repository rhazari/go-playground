package observer

import (
	"fmt"
)

type Observer interface {
    Update(message string)
}

type ConcreteObserver struct {
    ID string
}

func (o *ConcreteObserver) Update(message string) {
    fmt.Printf("Observer %s received message: %s\n", o.ID, message)
}

type Subject struct {
    observers []Observer
}

func (s *Subject) Register(observer Observer) {
    s.observers = append(s.observers, observer)
}

func (s *Subject) Deregister(observer Observer) {
    for i, obs := range s.observers {
        if obs == observer {
            // remove observer
            s.observers = append(s.observers[:i], s.observers[i+1:]...)
            break
        }
    }
}

func (s *Subject) NotifyAll(message string) {
    for _, observer := range s.observers {
        observer.Update(message)
    }
}