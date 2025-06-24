package main

import (
	"fmt"
	"time"
)

/*

Observer Design pattern allow some objects to notify other object about changes to thier state
its publisher and subscriber design pattern
its widely used by in asychronous message communication where events are sent from publisher to subscriber
if one object is related or depend to other object the if its changes the it should notify to other
*/

type Observer struct {
	name         string
	notification chan string
}

type Subject interface {
	Register(observer Observer)
	NotifyAll(msg string)
}

type NewAgency struct {
	observer []*Observer
	msg      string
}

func (a *NewAgency) Register(observer Observer) {
	a.observer = append(a.observer, &observer)
}

func (a *NewAgency) NotifyAll(msg string) {
	for _, o := range a.observer {
		m := fmt.Sprintf("Message %v sent to the observer named as: % v ", msg, o.name)
		o.notification <- m
	}
}

func ObserverPattern() {
	sub := &NewAgency{}

	o1 := Observer{
		name:         "Time of india",
		notification: make(chan string),
	}

	o2 := Observer{
		name:         "Deccan Herarld",
		notification: make(chan string),
	}

	sub.Register(o1)
	sub.Register(o2)

	go sub.NotifyAll("Brealing News !!")

	for {
		select {
		case msg := <-o1.notification:
			fmt.Println("Recived message on o1", msg)
		case msg2 := <-o2.notification:
			fmt.Println("Recived message on o2", msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("No messgae on any observer recivved ")
		}

	}
}
