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
	_, err := tea.NewProgram(sw.NewStacker(&Login{})).Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully logged in!")
}
