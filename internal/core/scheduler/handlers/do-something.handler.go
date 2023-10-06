package handlers

import "log"

func (hs *Storage) DoSomethingHandler([]HandlerParam) func(interface{}) bool {
	return func(interface{}) bool {
		log.Printf("Do something .....")
		return true
	}
}
