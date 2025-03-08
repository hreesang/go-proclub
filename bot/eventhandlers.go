package bot

import (
	"reflect"
	"sync"
	"time"
)

var (
	eventHandlers   = make([]any, 0)
	eventWg         sync.WaitGroup
)

func registerEventHandlers() {
	for _, handler := range eventHandlers {
		Session.AddHandler(handler)
	}
}

func wrapHandler(handler any) any {
	fnType := reflect.TypeOf(handler)

	wrappedHandler := reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		eventWg.Add(1)
		defer eventWg.Done()
		
		return reflect.ValueOf(handler).Call(args)
	})

	return wrappedHandler.Interface()
}

func AddEventHandler(handler any) {
	wrappedHandler := wrapHandler(handler)

	eventHandlers = append(eventHandlers, wrappedHandler)
}

func WaitEventHandlers() {
	time.Sleep(time.Second * 1)
	eventWg.Wait()
}
