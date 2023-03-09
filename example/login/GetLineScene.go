package main

import (
	"fmt"

	sw "github.com/CalebJohnHunt/stacker"

	tea "github.com/charmbracelet/bubbletea"
)

type GetLineScene struct {
	prompt string
	line   []rune
}

func (g *GetLineScene) Init() tea.Cmd {
	return nil
}

func (g *GetLineScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var batch tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "enter":
			return g, sw.PopScene(string(g.line))
		case "backspace":
			g.line = g.line[:len(g.line)-1]
		default:
			if len(key) != 1 {
				break
			}
			g.line = append(g.line, rune(key[0]))
		}
	}
	return g, batch
}

func (g *GetLineScene) View() string {
	return fmt.Sprintf("%s\n%s", g.prompt, string(g.line))
}
