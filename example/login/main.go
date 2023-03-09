package main

import (
	"fmt"
	"time"

	sw "github.com/CalebJohnHunt/stacker"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Println("Let's log in!")
	time.Sleep(time.Second)
	m, err := tea.NewProgram(sw.NewSwitcher(&Login{})).Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully logged in! And got auth token: %s\n", m.(*sw.Stacker).Peek().(*Login).authToken)
}
