package observer

import "testing"

func Test_Observer(t *testing.T) {
	sub := &Subject{}

    obs1 := &ConcreteObserver{ID: "A"}
    obs2 := &ConcreteObserver{ID: "B"}

    sub.Register(obs1)
    sub.Register(obs2)

    sub.NotifyAll("First message")
    sub.Deregister(obs1)
    sub.NotifyAll("Second message")
}