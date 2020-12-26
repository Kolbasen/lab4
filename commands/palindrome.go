package commands

import (
	"github.com/Kolbasen/lab4/engine"
)

// PalindromeCommand - represends command which can return palindrom of input
type PalindromeCommand struct {
	Arg string
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func isPalindrome(str string) bool {
	return str == reverse(str)
}

// Execute - check if input is already palindrom. If yes - returns input, if no - creates a palindrom.
func (p *PalindromeCommand) Execute(loop engine.Handler) {
	if isPalindrome(p.Arg) {
		loop.Post(&PrintCommand{Arg: p.Arg})
		return
	}
	reversedString := reverse(p.Arg)
	palindrom := p.Arg + reversedString
	loop.Post(&PrintCommand{Arg: palindrom})
	return
}
