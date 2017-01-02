package bank

import (
	"github.com/mishudark/eventhus"
	"log"
)

type CommandHandler struct {
	repository *eventhus.Repository
}

func NewCommandHandler(repository *eventhus.Repository) *CommandHandler {
	return &CommandHandler{
		repository: repository,
	}
}

func (c *CommandHandler) Handle(command eventhus.Command) error {
	var err error
	var version int
	var account Account

	switch command.(type) {
	case CreateAccount:
		version = 0

	default:
		if err = c.repository.Load(&account, command.GetAggregateID()); err != nil {
			log.Println(err)
			return err
		}

		version = command.GetVersion()
	}

	if err = account.Handle(command); err != nil {
		return err
	}

	if err = c.repository.Save(&account, version); err != nil {
		log.Println(err)
		return err
	}

	if err = c.repository.PublishEvents(&account, "bank", "account"); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
