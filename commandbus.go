package cqrs

//CommandBus serve as the bridge between commands and command handler
//it should manage the queues
type CommandBus interface {
	Add(command interface{})
	Dispatch(command interface{})
}
