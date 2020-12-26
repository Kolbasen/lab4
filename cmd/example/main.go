package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/Kolbasen/lab4/commands"
	"github.com/Kolbasen/lab4/engine"
)

func parse(commandline string) engine.Command {
	commandArgs := strings.Fields(commandline)
	commandName, args := commandArgs[0], commandArgs[1:]
	command, err := commands.GetCommand(commandName, args)
	if err != nil {
		command = &commands.PrintCommand{Arg: err.Error()}
		return command
	}
	return command
}

func main() {
	inputFile := flag.String("f", "", "File to read")
	flag.Parse()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open(*inputFile); err == nil {
		defer input.Close()

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			eventLoop.Post(cmd)
		}
		eventLoop.AwaitFinish()
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
