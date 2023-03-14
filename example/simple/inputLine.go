package main

import (
	"fmt"

	sw "github.com/CalebJohnHunt/stacker"

	tea "github.com/charmbracelet/bubbletea"
)

type inputLine struct {
	input  []rune
	silent bool
}

func (i *inputLine) Init() tea.Cmd {
	return nil
}

func (i *inputLine) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if i.silent {
				return i, sw.PopSceneSilent()
			}
			return i, sw.PopScene()
		case "backspace":
			i.input = i.input[:len(i.input)-1]
		default:
			if len(msg.String()) != 1 {
				break
			}
			i.input = append(i.input, rune(msg.String()[0]))
		}
	}
	return i, nil
}

func (i *inputLine) View() string {
	return fmt.Sprintf("Input:\n%s", string(i.input))
}
