package commands

import (
	"errors"
	"reflect"

	"github.com/Kolbasen/lab4/engine"
)

func stringToCommand(commandName string, arguments []string) (command engine.Command, err error) {
	switch commandName {
	case "print":
		return &PrintCommand{Arg: arguments[0]}, nil
	case "palindrome":
		return &PalindromeCommand{Arg: arguments[0]}, nil
	default:
		return nil, errors.New("Wrong command name")
	}
}

// GetCommand ...
func GetCommand(commandName string, arguments []string) (command engine.Command, err error) {
	command, err = stringToCommand(commandName, arguments)
	if err != nil {
		return nil, err
	}
	r := reflect.ValueOf(command)
	if len(arguments) != r.Elem().NumField() {
		return nil, errors.New("Wrong number of arguments")
	}
	return command, nil
}
